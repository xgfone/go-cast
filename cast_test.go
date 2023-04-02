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

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

// String is used to implement the interface fmt.Stringer.
type String string

// String implements the interface fmt.Stringer.
func (s String) String() string { return string(s) }

func ExampleToBool() {
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
	fmt.Println(ToBool([]byte{'\x31'})) // => string("1") => true
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
	// true <nil>
	// false <nil>
	// true <nil>
	// true <nil>
	// false <nil>
	// true <nil>
}

func ExampleToInt64() {
	fmt.Println(ToInt64(nil))
	fmt.Println(ToInt64(false))
	fmt.Println(ToInt64(true))
	fmt.Println(ToInt64(""))
	fmt.Println(ToInt64("123"))
	fmt.Println(ToInt64([]byte("456")))
	fmt.Println(ToInt64(789.0))
	fmt.Println(ToInt64(100))
	fmt.Println(ToInt64(time.Second)) // => int64(ms)
	fmt.Println(ToInt64(time.Unix(1234567890, 0)))
	fmt.Println(ToInt64(String("123456789"))) // Use the method String()

	// Output:
	// 0 <nil>
	// 0 <nil>
	// 1 <nil>
	// 0 <nil>
	// 123 <nil>
	// 456 <nil>
	// 789 <nil>
	// 100 <nil>
	// 1000 <nil>
	// 1234567890 <nil>
	// 123456789 <nil>
}

func ExampleToUint64() {
	fmt.Println(ToUint64(nil))
	fmt.Println(ToUint64(false))
	fmt.Println(ToUint64(true))
	fmt.Println(ToUint64(""))
	fmt.Println(ToUint64("123"))
	fmt.Println(ToUint64([]byte("456")))
	fmt.Println(ToUint64(789.0))
	fmt.Println(ToUint64(100))
	fmt.Println(ToUint64(String("123456789"))) // Use the method String()
	// fmt.Println(ToUint64(time.Second))              // unsupported
	// fmt.Println(ToUint64(time.Unix(1234567890, 0))) // unsupported

	// Output:
	// 0 <nil>
	// 0 <nil>
	// 1 <nil>
	// 0 <nil>
	// 123 <nil>
	// 456 <nil>
	// 789 <nil>
	// 100 <nil>
	// 123456789 <nil>
}

func ExampleToFloat64() {
	fmt.Println(ToFloat64(nil))
	fmt.Println(ToFloat64(false))
	fmt.Println(ToFloat64(true))
	fmt.Println(ToFloat64(""))
	fmt.Println(ToFloat64("123"))
	fmt.Println(ToFloat64([]byte("456")))
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
	// 456 <nil>
	// 789 <nil>
	// 100 <nil>
	// 200 <nil>
	// 1 <nil>
}

func ExampleToString() {
	fmt.Println(ToString(nil))
	fmt.Println(ToString(false))
	fmt.Println(ToString(true))
	fmt.Println(ToString("123"))
	fmt.Println(ToString([]byte("456")))
	fmt.Println(ToString(789.0))
	fmt.Println(ToString(100))
	fmt.Println(ToString(String("123456789")))                   // Use the method String()
	fmt.Println(ToString(time.Second))                           // Use the method String()
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

func ExampleToDuration() {
	fmt.Println(ToDuration(nil))
	fmt.Println(ToDuration(""))     // Parse string as time.Millisecond
	fmt.Println(ToDuration("1000")) // Parse string as time.Millisecond
	fmt.Println(ToDuration("2s"))   // Use time.ParseDuration
	fmt.Println(ToDuration([]byte("3000")))
	fmt.Println(ToDuration([]byte("4s")))
	fmt.Println(ToDuration(5.0))  // Use float as time.Second
	fmt.Println(ToDuration(6000)) // Use integer as time.Millisecond
	fmt.Println(ToDuration(time.Second))
	fmt.Println(ToDuration(String("7000"))) // Use the method String() to be parsed as time.Millisecond
	fmt.Println(ToDuration(String("8s")))   // Use the method String() to be parsed by time.ParseDurationParseDuration

	// Output:
	// 0s <nil>
	// 0s <nil>
	// 1s <nil>
	// 2s <nil>
	// 3s <nil>
	// 4s <nil>
	// 5s <nil>
	// 6s <nil>
	// 1s <nil>
	// 7s <nil>
	// 8s <nil>
}

func ExampleToTime() {
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

func TestIsIntegerString(t *testing.T) {
	const s = "-1000"
	if !isIntegerString(s) {
		t.Errorf("%s: expect true, but got false", s)
	} else if v, err := strconv.ParseInt(s, 10, 64); err != nil {
		t.Error(err)
	} else if v != -1000 {
		t.Errorf("expect %d, but got %d", -1000, v)
	}
}
