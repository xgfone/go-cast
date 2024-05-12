// Copyright 2023 xgfone
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

package cast

import "time"

// Must checks whether err is not nil. If not nil, panic with it. Or, return value.
func Must[T any](value T, err error) T {
	if err == nil {
		return value
	}
	panic(err)
}

// MustToTimeInLocation is the same as ToTimeInLocation, but panics if there is an error.
func MustToTimeInLocation(any interface{}, loc *time.Location, layouts ...string) time.Time {
	return Must(ToTimeInLocation(any, loc, layouts...))
}

// MustParseTime is the same as TryParseTime, but panics if there is an error.
func MustParseTime(value string, loc *time.Location, layouts ...string) time.Time {
	return Must(TryParseTime(value, loc, layouts...))
}
