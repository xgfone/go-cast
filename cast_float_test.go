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

func ExampleToFloat64() {
	type String string

	fmt.Println(ToFloat64(nil))
	fmt.Println(ToFloat64(false))
	fmt.Println(ToFloat64(true))
	fmt.Println(ToFloat64(""))
	fmt.Println(ToFloat64("123"))
	fmt.Println(ToFloat64(789.0))
	fmt.Println(ToFloat64(100))
	fmt.Println(ToFloat64(String("200.0"))) // Use the method String()
	fmt.Println(ToFloat64(time.Second))     // N<s> => float64(N)
	// fmt.Println(ToFloat64(time.Unix(1234567890, 0))) // unsupported

	// Output:
	// 0 <nil>
	// 0 <nil>
	// 1 <nil>
	// 0 <nil>
	// 123 <nil>
	// 789 <nil>
	// 100 <nil>
	// 200 <nil>
	// 1 <nil>
}
