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

func ExampleToDuration() {
	type String string

	fmt.Println(ToDuration(nil))
	fmt.Println(ToDuration(""))   // Parse string as time.Millisecond
	fmt.Println(ToDuration("2s")) // Use time.ParseDuration
	fmt.Println(ToDuration(5.0))  // Use float as time.Second
	fmt.Println(ToDuration(6000)) // Use integer as time.Millisecond
	fmt.Println(ToDuration(time.Second))
	fmt.Println(ToDuration(String("8s"))) // Use the method String() to be parsed by time.ParseDurationParseDuration

	// Output:
	// 0s <nil>
	// 0s <nil>
	// 2s <nil>
	// 5s <nil>
	// 6s <nil>
	// 1s <nil>
	// 8s <nil>
}
