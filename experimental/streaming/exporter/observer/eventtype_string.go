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

// Code generated by "stringer -type=EventType"; DO NOT EDIT.

package observer

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[INVALID-0]
	_ = x[START_SPAN-1]
	_ = x[FINISH_SPAN-2]
	_ = x[ADD_EVENT-3]
	_ = x[ADD_EVENTF-4]
	_ = x[NEW_SCOPE-5]
	_ = x[NEW_MEASURE-6]
	_ = x[NEW_METRIC-7]
	_ = x[MODIFY_ATTR-8]
	_ = x[RECORD_STATS-9]
}

const _EventType_name = "INVALIDSTART_SPANFINISH_SPANLOG_EVENTLOGF_EVENTNEW_SCOPENEW_MEASURENEW_METRICMODIFY_ATTRRECORD_STATS"

var _EventType_index = [...]uint8{0, 7, 17, 28, 37, 47, 56, 67, 77, 88, 100}

func (i EventType) String() string {
	if i < 0 || i >= EventType(len(_EventType_index)-1) {
		return "EventType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _EventType_name[_EventType_index[i]:_EventType_index[i+1]]
}
