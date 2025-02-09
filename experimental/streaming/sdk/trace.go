// Copyright 2019, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sdk

import (
	"context"
	"math/rand"

	"go.opentelemetry.io/api/core"
	"go.opentelemetry.io/api/key"
	"go.opentelemetry.io/api/tag"
	"go.opentelemetry.io/api/trace"
	apitrace "go.opentelemetry.io/api/trace"
	"go.opentelemetry.io/experimental/streaming/exporter/observer"
)

type tracer struct {
	resources observer.EventID
}

var (
	// TODO These should move somewhere in the api, right?
	ServiceKey   = key.New("service")
	ComponentKey = key.New("component")
	ErrorKey     = key.New("error")
	SpanIDKey    = key.New("span_id")
	TraceIDKey   = key.New("trace_id")
	MessageKey   = key.New("message",
		key.WithDescription("message text: info, error, etc"),
	)
)

func New() trace.Tracer {
	return &tracer{}
}

func (t *tracer) WithResources(attributes ...core.KeyValue) apitrace.Tracer {
	s := observer.NewScope(observer.ScopeID{
		EventID: t.resources,
	}, attributes...)
	return &tracer{
		resources: s.EventID,
	}
}

func (t *tracer) WithComponent(name string) apitrace.Tracer {
	return t.WithResources(ComponentKey.String(name))
}

func (t *tracer) WithService(name string) apitrace.Tracer {
	return t.WithResources(ServiceKey.String(name))
}

func (t *tracer) WithSpan(ctx context.Context, name string, body func(context.Context) error) error {
	// TODO: use runtime/trace.WithRegion for execution tracer support
	// TODO: use runtime/pprof.Do for profile tags support
	ctx, span := t.Start(ctx, name)
	defer span.Finish()

	if err := body(ctx); err != nil {
		span.SetAttribute(ErrorKey.Bool(true))
		span.Event(ctx, "span error", MessageKey.String(err.Error()))
		return err
	}
	return nil
}

func (t *tracer) Start(ctx context.Context, name string, opts ...apitrace.SpanOption) (context.Context, apitrace.Span) {
	var child core.SpanContext

	child.SpanID = rand.Uint64()

	o := &apitrace.SpanOptions{}

	for _, opt := range opts {
		opt(o)
	}

	var parentScope observer.ScopeID

	if o.Reference.HasTraceID() {
		parentScope.SpanContext = o.Reference.SpanContext
	} else {
		parentScope.SpanContext = apitrace.CurrentSpan(ctx).SpanContext()
	}

	if parentScope.HasTraceID() {
		parent := parentScope.SpanContext
		child.TraceID.High = parent.TraceID.High
		child.TraceID.Low = parent.TraceID.Low
	} else {
		child.TraceID.High = rand.Uint64()
		child.TraceID.Low = rand.Uint64()
	}

	childScope := observer.ScopeID{
		SpanContext: child,
		EventID:     t.resources,
	}

	span := &span{
		tracer: t,
		initial: observer.ScopeID{
			SpanContext: child,
			EventID: observer.Record(observer.Event{
				Time:    o.StartTime,
				Type:    observer.START_SPAN,
				Scope:   observer.NewScope(childScope, o.Attributes...),
				Context: ctx,
				Parent:  parentScope,
				String:  name,
			},
			),
		},
	}
	return trace.SetCurrentSpan(ctx, span), span
}

func (t *tracer) Inject(ctx context.Context, span apitrace.Span, injector apitrace.Injector) {
	injector.Inject(span.SpanContext(), tag.FromContext(ctx))
}
