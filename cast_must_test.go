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

func ExampleMust() {
	Must(ToBool("true"))
	Must(ToInt64("123"))
	Must(ToString(1234))
	// Output:
}

func ExampleMustParseTime() {
	MustParseTime("2009-02-13T23:31:30Z", nil)
	MustParseTime("2009-02-13 23:31:30", time.UTC)
	// Output:
}

func ExampleMustToTimeInLocation() {
	MustToTimeInLocation("2009-02-13T23:31:30Z", nil)
	MustToTimeInLocation(1234567890, nil)
	// Output:
}
