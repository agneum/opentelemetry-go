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

package stdout

import (
	"os"

	"go.opentelemetry.io/experimental/streaming/exporter/observer"
	"go.opentelemetry.io/experimental/streaming/exporter/reader"
	"go.opentelemetry.io/experimental/streaming/exporter/reader/format"
)

type stdoutLog struct{}

func New() observer.Observer {
	return reader.NewReaderObserver(&stdoutLog{})
}

func (s *stdoutLog) Read(data reader.Event) {
	os.Stdout.WriteString(format.EventToString(data))
}
