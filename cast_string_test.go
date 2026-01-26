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

func ExampleToString() {
	type String string

	fmt.Println(ToString(nil))
	fmt.Println(ToString(false))
	fmt.Println(ToString(true))
	fmt.Println(ToString("123"))
	fmt.Println(ToString([]byte("456")))
	fmt.Println(ToString(789.0))
	fmt.Println(ToString(100))
	fmt.Println(ToString(String("123456789")))
	fmt.Println(ToString(time.Second))
	fmt.Println(ToString(time.Unix(1234567890, 0).In(time.UTC))) // Use time.RFC3339Nano

	// Output:
	//  <nil>
	// false <nil>
	// true <nil>
	// 123 <nil>
	// 456 <nil>
	// 789 <nil>
	// 100 <nil>
	// 123456789 <nil>
	// 1s <nil>
	// 2009-02-13T23:31:30Z <nil>
}
