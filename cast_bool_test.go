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

func ExampleToBool() {
	type String string

	fmt.Println(ToBool(nil))
	fmt.Println(ToBool(false))
	fmt.Println(ToBool(true))
	fmt.Println(ToBool(""))
	fmt.Println(ToBool("0"))
	fmt.Println(ToBool("1"))
	fmt.Println(ToBool("false"))
	fmt.Println(ToBool("true"))
	fmt.Println(ToBool([]byte{'\x00'})) // => byte(0)     => false
	fmt.Println(ToBool([]byte{'\x01'})) // => byte(1)     => true
	fmt.Println(ToBool(0))
	fmt.Println(ToBool(1))
	fmt.Println(ToBool(2))
	fmt.Println(ToBool(String("f"))) // Use the method String()
	fmt.Println(ToBool(String("t"))) // Use the method String()

	// Output:
	// false <nil>
	// false <nil>
	// true <nil>
	// false <nil>
	// false <nil>
	// true <nil>
	// false <nil>
	// true <nil>
	// false <nil>
	// true <nil>
	// false <nil>
	// true <nil>
	// true <nil>
	// false <nil>
	// true <nil>
}
