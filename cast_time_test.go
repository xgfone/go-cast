// Copyright 2026 xgfone
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

import (
	"fmt"
	"time"
)

func ExampleToTime() {
	type String string

	fmt.Println(ToTime(nil))
	fmt.Println(ToTime(""))                                   // Use the method String() to be parsed by time.Parse
	fmt.Println(ToTime("0000-00-00 00:00:00"))                // Use the method String() to be parsed by time.Parse
	fmt.Println(ToTime("0000-00-00 00:00:00.000"))            // Use the method String() to be parsed by time.Parse
	fmt.Println(ToTime("0000-00-00 00:00:00.000000"))         // Use the method String() to be parsed by time.Parse
	fmt.Println(ToTime("2009-02-13 23:31:30"))                // Use the method String() to be parsed by time.Parse
	fmt.Println(ToTime("2009-02-13T23:31:30Z"))               // Use the method String() to be parsed by time.Parse
	fmt.Println(ToTime([]byte{}))                             // Use the method String() to be parsed by time.Parse
	fmt.Println(ToTime([]byte("0000-00-00 00:00:00")))        // Use the method String() to be parsed by time.Parse
	fmt.Println(ToTime([]byte("0000-00-00 00:00:00.000")))    // Use the method String() to be parsed by time.Parse
	fmt.Println(ToTime([]byte("0000-00-00 00:00:00.000000"))) // Use the method String() to be parsed by time.Parse
	fmt.Println(ToTime([]byte("2009-02-13 23:31:30")))        // Use the method String() to be parsed by time.Parse
	fmt.Println(ToTime([]byte("2009-02-13T23:31:30Z")))       // Use the method String() to be parsed by time.Parse
	fmt.Println(ToTime(String("2009-02-13T23:31:30Z")))       // Use the method String() to be parsed by time.Parse
	fmt.Println(ToTime(time.Unix(1234567890, 0)))
	fmt.Println(ToTime(1234567890.0)) // Use float as the unix timestamp
	fmt.Println(ToTime(1234567890))   // Use float as the unix timestamp

	// Output:
	// 0001-01-01 00:00:00 +0000 UTC <nil>
	// 0001-01-01 00:00:00 +0000 UTC <nil>
	// 0001-01-01 00:00:00 +0000 UTC <nil>
	// 0001-01-01 00:00:00 +0000 UTC <nil>
	// 0001-01-01 00:00:00 +0000 UTC <nil>
	// 2009-02-13 23:31:30 +0000 UTC <nil>
	// 2009-02-13 23:31:30 +0000 UTC <nil>
	// 0001-01-01 00:00:00 +0000 UTC <nil>
	// 0001-01-01 00:00:00 +0000 UTC <nil>
	// 0001-01-01 00:00:00 +0000 UTC <nil>
	// 0001-01-01 00:00:00 +0000 UTC <nil>
	// 2009-02-13 23:31:30 +0000 UTC <nil>
	// 2009-02-13 23:31:30 +0000 UTC <nil>
	// 2009-02-13 23:31:30 +0000 UTC <nil>
	// 2009-02-13 23:31:30 +0000 UTC <nil>
	// 2009-02-13 23:31:30 +0000 UTC <nil>
	// 2009-02-13 23:31:30 +0000 UTC <nil>
}
