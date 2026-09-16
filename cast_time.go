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
	"reflect"
	"strconv"
	"time"

	"github.com/xgfone/go-toolkit/timex"
	"github.com/xgfone/go-toolkit/unsafex"
)

var timeFormats = []string{
	time.RFC3339Nano,
	"2006-01-02 15:04:05",
	"2006-01-02",
}

// SetTimeLocation forwards the calling to [github.com/xgfone/go-toolkit/timex.SetLocation].
func SetTimeLocation(loc *time.Location) {
	timex.SetLocation(loc)
}

// SetTimeFormats sets the default time formats used by TryParseTime.
//
// Default: [RFC3339Nano, "2006-01-02 15:04:05", "2006-01-02"]
func SetTimeFormats(formats ...string) {
	if len(formats) == 0 {
		panic("cast.SetTimeFormats: formats is empty")
	}
	for i, format := range formats {
		if format == "" {
			panic(fmt.Errorf("cast.SetTimeFormats: %dth format is empty", i))
		}
	}
	timeFormats = formats
}

// ToTime is a convenient function, which is equal to ToTimeInLocation(any, nil).
func ToTime(any any) (dst time.Time, err error) {
	return ToTimeInLocation(any, nil)
}

// ToTimeInLocation converts any to a time.Time value.
//
// Supports the types as follow:
//
//	~string: => TryParseTime
//	~float32, ~float64: => unix timestamp
//	~int, ~int8, ~int16, ~int32, ~int64: => unix timestamp
//	~uint, ~uint8, ~uint16, ~uint32, ~uint64: => unix timestamp
//	time.Time
//
// And the pointer to types above, and the types as follow:
//
//	nil
//	[]byte
//	interface{ Time() time.Time }
//
// If loc is nil, use timex.Location instead.
// If any is a string-like, use TryParseTime to parse it with layouts.
func ToTimeInLocation(any any, loc *time.Location) (dst time.Time, err error) {
	if loc == nil {
		loc = getLocation()
	}

	switch src := any.(type) {
	case nil:
		dst = dst.In(loc)
	case string:
		dst, err = TryParseTime(src, loc)
	case []byte:
		dst, err = TryParseTime(unsafex.String(src), loc)
	case float32:
		dst = time.Unix(int64(src), 0).In(loc)
	case float64:
		dst = time.Unix(int64(src), 0).In(loc)
	case int:
		dst = time.Unix(int64(src), 0).In(loc)
	case int32:
		dst = time.Unix(int64(src), 0).In(loc)
	case int64:
		dst = time.Unix(int64(src), 0).In(loc)
	case uint:
		dst = time.Unix(int64(src), 0).In(loc)
	case uint32:
		dst = time.Unix(int64(src), 0).In(loc)
	case uint64:
		dst = time.Unix(int64(src), 0).In(loc)
	case time.Time:
		dst = src.In(loc)
	case *time.Time:
		dst = src.In(loc)
	case interface{ Time() time.Time }:
		dst = src.Time()
	default:
		dst, err = tryReflectToTimeInLocation(reflect.ValueOf(any), loc)
	}

	return
}

func tryReflectToTimeInLocation(src reflect.Value, loc *time.Location, layouts ...string) (dst time.Time, err error) {
	switch src.Kind() {
	case reflect.Invalid:
	case reflect.Pointer:
		if !src.IsNil() {
			dst, err = tryReflectToTimeInLocation(src.Elem(), loc, layouts...)
		}

	case reflect.String:
		dst, err = TryParseTime(src.String(), loc, layouts...)

	case reflect.Float32, reflect.Float64:
		dst = time.Unix(int64(src.Float()), 0).In(loc)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		dst = time.Unix(src.Int(), 0).In(loc)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		dst = time.Unix(int64(src.Uint()), 0).In(loc)

	default:
		err = fmt.Errorf("cast.ToTimeInLocation: unsupport to convert %T to time.Time", src.Interface())
	}

	return
}

// TryParseTime tries to parse the string value with the layouts in turn to time.Time.
//
// If loc is nil, use timex.Location instead.
// If layouts is empty, use the default formats set by SetTimeFormats instead.
//
// If value is an integer string, it will be parsee as the unix timestamp.
func TryParseTime(value string, loc *time.Location, layouts ...string) (time.Time, error) {
	if loc == nil {
		loc = getLocation()
	}

	switch value {
	case
		"0000-00-00 00:00:00",
		"0000-00-00 00:00:00.000",
		"0000-00-00 00:00:00.000000",
		"0001-01-01 00:00:00",
		"0001-01-01 00:00:00.000",
		"0001-01-01 00:00:00.000000",
		"":
		return time.Time{}.In(loc), nil
	}

	if isIntegerString(value) {
		i, err := strconv.ParseInt(value, 10, 64)
		return time.Unix(i, 0).In(loc), err
	}

	if len(layouts) == 0 {
		layouts = timeFormats
		if len(layouts) == 0 {
			layouts = []string{time.DateTime, time.RFC3339Nano}
		}
	}

	for _, layout := range layouts {
		if t, err := time.ParseInLocation(layout, value, loc); err == nil {
			return t, nil
		}
	}

	return time.Time{}.In(loc), fmt.Errorf("unable to parse time '%s'", value)
}

func getLocation() *time.Location {
	switch loc := any(timex.Location).(type) {
	case *time.Location: //<1.0.0
		return loc

	case func() *time.Location: // >=v1.0.0
		return loc()

	default:
		panic(fmt.Errorf("unknown timex.Location type '%T'", loc))
	}
}
