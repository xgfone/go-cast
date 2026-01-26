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

import "fmt"

func ExampleToUint64() {
	type String string

	fmt.Println(ToUint64(nil))
	fmt.Println(ToUint64(false))
	fmt.Println(ToUint64(true))
	fmt.Println(ToUint64(""))
	fmt.Println(ToUint64("123"))
	fmt.Println(ToUint64(789.0))
	fmt.Println(ToUint64(100))
	fmt.Println(ToUint64(String("123456789")))
	fmt.Println(ToUint64([]byte("\x01")))

	// Output:
	// 0 <nil>
	// 0 <nil>
	// 1 <nil>
	// 0 <nil>
	// 123 <nil>
	// 789 <nil>
	// 100 <nil>
	// 123456789 <nil>
	// 1 <nil>
}
