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
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/xgfone/go-defaults"
)

// Define some hook functions to intercept the ToXXX conversion.
var (
	ToBoolHook     func(src interface{}) (dst bool, err error)
	ToInt64Hook    func(src interface{}) (dst int64, err error)
	ToUint64Hook   func(src interface{}) (dst uint64, err error)
	ToFloat64Hook  func(src interface{}) (dst float64, err error)
	ToStringHook   func(src interface{}) (dst string, err error)
	ToDurationHook func(src interface{}) (dst time.Duration, err error)
	ToTimeHook     func(src interface{}, loc *time.Location, layouts ...string) (dst time.Time, err error)
)

// ToBool prefers to use ToBoolHook to convert any to a bool value
// rather than ToBoolPure.
func ToBool(any interface{}) (dst bool, err error) {
	if ToBoolHook != nil {
		dst, err = ToBoolHook(any)
	} else {
		dst, err = ToBoolPure(any)
	}
	return
}

// ToInt64 prefers to use ToInt64Hook to convert any to a int64 value
// rather than ToInt64Pure.
func ToInt64(any interface{}) (dst int64, err error) {
	if ToInt64Hook != nil {
		dst, err = ToInt64Hook(any)
	} else {
		dst, err = ToInt64Pure(any)
	}
	return
}

// ToUint64 prefers to use ToUint64Hook to convert any to a uint64 value
// rather than ToUint64Pure.
func ToUint64(any interface{}) (dst uint64, err error) {
	if ToUint64Hook != nil {
		dst, err = ToUint64Hook(any)
	} else {
		dst, err = ToUint64Pure(any)
	}
	return
}

// ToFloat64 prefers to use ToFloat64Hook to convert any to a float64 value
// rather than ToFloat64Pure.
func ToFloat64(any interface{}) (dst float64, err error) {
	if ToFloat64Hook != nil {
		dst, err = ToFloat64Hook(any)
	} else {
		dst, err = ToFloat64Pure(any)
	}
	return
}

// ToString prefers to use ToStringHook to convert any to a string value
// rather than ToStringPure.
func ToString(any interface{}) (dst string, err error) {
	if ToStringHook != nil {
		dst, err = ToStringHook(any)
	} else {
		dst, err = ToStringPure(any)
	}
	return
}

// ToDuration prefers to use ToDurationHook to convert any to a time.Duration value
// rather than ToDurationPure.
func ToDuration(any interface{}) (dst time.Duration, err error) {
	if ToDurationHook != nil {
		dst, err = ToDurationHook(any)
	} else {
		dst, err = ToDurationPure(any)
	}
	return
}

// ToTimeInLocation prefers to use ToTimeHook to convert any to a time.Time value
// rather than ToTimeInLocationPure.
func ToTimeInLocation(any interface{}, loc *time.Location, layouts ...string) (dst time.Time, err error) {
	if ToTimeHook != nil {
		if len(layouts) == 0 {
			layouts = defaults.TimeFormats.Get()
		}
		dst, err = ToTimeHook(any, loc, layouts...)
	} else {
		dst, err = ToTimeInLocationPure(any, loc, layouts...)
	}
	return
}

// ToTime is a convenient function, which is equal to ToTimeInLocation(any, nil).
func ToTime(any interface{}) (dst time.Time, err error) {
	return ToTimeInLocation(any, nil)
}

// ToBoolPure converts any to a bool value.
//
// Supports the types as follow:
//
//	~bool
//	~string: => strconv.ParseBool
//	~float32, ~float64: => !=0
//	~int, ~int8, ~int16, ~int32, ~int64: => !=0
//	~uint, ~uint8, ~uint16, ~uint32, ~uint64, ~uintptr: => !=0
//
// And the pointer to types above, and the types as follow:
//
//	nil
//	[]byte
//	fmt.Stringer
//	interface{ Bool() bool }
//	interface{ IsZero() bool }
func ToBoolPure(any interface{}) (dst bool, err error) {
	switch src := any.(type) {
	case nil:
	case bool:
		dst = src
	case string:
		dst, err = parseBool(src)
	case []byte:
		switch len(src) {
		case 0:
		case 1:
			switch src[0] {
			case '\x00':
			case '\x01':
				dst = true
			default:
				dst, err = parseBool(string(src))
			}
		default:
			dst, err = parseBool(string(src))
		}
	case float32:
		dst = src != 0
	case float64:
		dst = src != 0
	case int:
		dst = src != 0
	case int8:
		dst = src != 0
	case int16:
		dst = src != 0
	case int32:
		dst = src != 0
	case int64:
		dst = src != 0
	case uint:
		dst = src != 0
	case uint8:
		dst = src != 0
	case uint16:
		dst = src != 0
	case uint32:
		dst = src != 0
	case uint64:
		dst = src != 0
	case uintptr:
		dst = src != 0
	case interface{ Bool() bool }:
		dst = src.Bool()
	case interface{ IsZero() bool }:
		dst = !src.IsZero()
	case fmt.Stringer:
		dst, err = parseBool(src.String())
	default:
		dst, err = tryReflectToBool(reflect.ValueOf(any))
	}
	return
}

func tryReflectToBool(src reflect.Value) (dst bool, err error) {
	switch src.Kind() {
	case reflect.Invalid:
	case reflect.Pointer:
		if !src.IsNil() {
			dst, err = tryReflectToBool(src.Elem())
		}

	case reflect.Bool:
		dst = src.Bool()

	case reflect.String:
		dst, err = parseBool(src.String())

	case reflect.Float32, reflect.Float64:
		dst = src.Float() != 0

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		dst = src.Int() != 0

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		dst = src.Uint() != 0

	default:
		err = fmt.Errorf("cast.ToBool: unsupport to convert %T to bool", src.Interface())
	}

	return
}

func parseBool(src string) (dst bool, err error) {
	if src != "" {
		dst, err = strconv.ParseBool(src)
	}
	return
}

// ToStringPure converts any to a string value.
//
// Supports the types as follow:
//
//	~bool
//	~string
//	~float32, ~float64
//	~int, ~int8, ~int16, ~int32, ~int64
//	~uint, ~uint8, ~uint16, ~uint32, ~uint64, ~uintptr
//	time.Time: => time.RFC3339Nano
//
// And the pointer to types above, and the types as follow:
//
//	nil
//	[]byte
//	error
//	fmt.Stringer
func ToStringPure(any interface{}) (dst string, err error) {
	switch src := any.(type) {
	case nil:
	case bool:
		dst = strconv.FormatBool(src)
	case string:
		dst = src
	case []byte:
		dst = string(src)
	case float32:
		dst = strconv.FormatFloat(float64(src), 'f', -1, 32)
	case float64:
		dst = strconv.FormatFloat(src, 'f', -1, 64)
	case int:
		dst = strconv.FormatInt(int64(src), 10)
	case int8:
		dst = strconv.FormatInt(int64(src), 10)
	case int16:
		dst = strconv.FormatInt(int64(src), 10)
	case int32:
		dst = strconv.FormatInt(int64(src), 10)
	case int64:
		dst = strconv.FormatInt(src, 10)
	case uint:
		return strconv.FormatUint(uint64(src), 10), nil
	case uint8:
		return strconv.FormatUint(uint64(src), 10), nil
	case uint16:
		return strconv.FormatUint(uint64(src), 10), nil
	case uint32:
		return strconv.FormatUint(uint64(src), 10), nil
	case uint64:
		return strconv.FormatUint(src, 10), nil
	case uintptr:
		return strconv.FormatUint(uint64(src), 10), nil
	case time.Time:
		dst = src.Format(time.RFC3339Nano)
	case *time.Time:
		dst = src.Format(time.RFC3339Nano)
	case error:
		dst = src.Error()
	case fmt.Stringer:
		dst = src.String()
	default:
		dst, err = tryReflectToString(reflect.ValueOf(any))
	}
	return
}

func tryReflectToString(src reflect.Value) (dst string, err error) {
	switch src.Kind() {
	case reflect.Invalid:
	case reflect.Pointer:
		if !src.IsNil() {
			dst, err = tryReflectToString(src.Elem())
		}

	case reflect.Bool:
		dst = strconv.FormatBool(src.Bool())

	case reflect.String:
		dst = src.String()

	case reflect.Float32:
		dst = strconv.FormatFloat(src.Float(), 'f', -1, 32)

	case reflect.Float64:
		dst = strconv.FormatFloat(src.Float(), 'f', -1, 64)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		dst = strconv.FormatInt(src.Int(), 10)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		dst = strconv.FormatUint(src.Uint(), 10)

	default:
		err = fmt.Errorf("cast.ToString: unsupport to convert %T to string", src.Interface())
	}
	return
}

// ToInt64Pure converts any to a int64 value.
//
// Supports the types as follow:
//
//	~bool
//	~string: => strconv.ParseInt
//	~float32, ~float64
//	~int, ~int8, ~int16, ~int32, ~int64
//	~uint, ~uint8, ~uint16, ~uint32, ~uint64, ~uintptr
//	time.Duration: => N(ms)
//	time.Time: => unix timestamp
//
// And the pointer to types above, and the types as follow:
//
//	nil
//	[]byte
//	fmt.Stringer
//	interface{ Int64() int64 }
//	interface{ Int() int64 }
func ToInt64Pure(any interface{}) (dst int64, err error) {
	switch src := any.(type) {
	case nil:
	case bool:
		if src {
			dst = 1
		}
	case string:
		dst, err = parseInt64(src)
	case []byte:
		dst, err = parseInt64(string(src))
	case float32:
		dst = int64(src)
	case float64:
		dst = int64(src)
	case int:
		dst = int64(src)
	case int8:
		dst = int64(src)
	case int16:
		dst = int64(src)
	case int32:
		dst = int64(src)
	case int64:
		dst = src
	case uint:
		dst = int64(src)
	case uint8:
		dst = int64(src)
	case uint16:
		dst = int64(src)
	case uint32:
		dst = int64(src)
	case uint64:
		dst = int64(src)
	case uintptr:
		dst = int64(src)
	case time.Duration:
		dst = int64(src / time.Millisecond)
	case *time.Duration:
		dst = int64(*src / time.Millisecond)
	case time.Time:
		dst = src.Unix()
	case *time.Time:
		dst = src.Unix()
	case interface{ Int64() int64 }:
		dst = src.Int64()
	case interface{ Int() int64 }:
		dst = src.Int()
	case fmt.Stringer:
		dst, err = parseInt64(src.String())
	default:
		dst, err = tryReflectToInt64(reflect.ValueOf(any))
	}
	return
}

func tryReflectToInt64(src reflect.Value) (dst int64, err error) {
	switch src.Kind() {
	case reflect.Invalid:
	case reflect.Pointer:
		if !src.IsNil() {
			dst, err = tryReflectToInt64(src.Elem())
		}

	case reflect.Bool:
		if src.Bool() {
			dst = 1
		}

	case reflect.String:
		dst, err = parseInt64(src.String())

	case reflect.Float32, reflect.Float64:
		dst = int64(src.Float())

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		dst = src.Int()

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		dst = int64(src.Uint())

	default:
		err = fmt.Errorf("cast.ToInt64: unsupport to convert %T to int64", src.Interface())
	}

	return
}

func parseInt64(src string) (dst int64, err error) {
	if src != "" {
		dst, err = strconv.ParseInt(src, 0, 64)
	}
	return
}

// ToUint64Pure converts any to a uint64 value.
//
// Supports the types as follow:
//
//	~bool
//	~string: => strconv.ParseUint
//	~float32, ~float64
//	~int, ~int8, ~int16, ~int32, ~int64
//	~uint, ~uint8, ~uint16, ~uint32, ~uint64, ~uintptr
//
// And the pointer to types above, and the types as follow:
//
//	nil
//	[]byte
//	fmt.Stringer
//	interface{ Uint64() uint64 }
//	interface{ Uint() uint64 }
func ToUint64Pure(any interface{}) (dst uint64, err error) {
	switch src := any.(type) {
	case nil:
	case bool:
		if src {
			dst = 1
		}
	case string:
		dst, err = parseUint64(src)
	case []byte:
		dst, err = parseUint64(string(src))
	case float32:
		if src < 0 {
			return 0, errors.New("cannot convert a negative to uint64")
		}
		dst = uint64(src)
	case float64:
		if src < 0 {
			return 0, errors.New("cannot convert a negative to uint64")
		}
		dst = uint64(src)
	case int:
		if src < 0 {
			return 0, errors.New("cannot convert a negative to uint64")
		}
		dst = uint64(src)
	case int8:
		if src < 0 {
			return 0, errors.New("cannot convert a negative to uint64")
		}
		dst = uint64(src)
	case int16:
		if src < 0 {
			return 0, errors.New("cannot convert a negative to uint64")
		}
		dst = uint64(src)
	case int32:
		if src < 0 {
			return 0, errors.New("cannot convert a negative to uint64")
		}
		dst = uint64(src)
	case int64:
		if src < 0 {
			return 0, errors.New("cannot convert a negative to uint64")
		}
		dst = uint64(src)
	case uint:
		dst = uint64(src)
	case uint8:
		dst = uint64(src)
	case uint16:
		dst = uint64(src)
	case uint32:
		dst = uint64(src)
	case uint64:
		dst = src
	case uintptr:
		dst = uint64(src)
	case interface{ Uint64() uint64 }:
		dst = src.Uint64()
	case interface{ Uint() uint64 }:
		dst = src.Uint()
	case fmt.Stringer:
		dst, err = parseUint64(src.String())
	default:
		dst, err = tryReflectToUint64(reflect.ValueOf(any))
	}
	return
}

func tryReflectToUint64(src reflect.Value) (dst uint64, err error) {
	switch src.Kind() {
	case reflect.Invalid:
	case reflect.Pointer:
		if !src.IsNil() {
			dst, err = tryReflectToUint64(src.Elem())
		}

	case reflect.Bool:
		if src.Bool() {
			dst = 1
		}

	case reflect.String:
		dst, err = parseUint64(src.String())

	case reflect.Float32, reflect.Float64:
		if v := src.Float(); v < 0 {
			err = errors.New("cannot convert a negative float to uint64")
		} else {
			dst = uint64(v)
		}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if v := src.Int(); v < 0 {
			err = errors.New("cannot convert a negative integer to uint64")
		} else {
			dst = uint64(v)
		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		dst = src.Uint()

	default:
		err = fmt.Errorf("cast.ToUint64: unsupport to convert %T to uint64", src.Interface())
	}

	return
}

func parseUint64(src string) (dst uint64, err error) {
	if src != "" {
		dst, err = strconv.ParseUint(src, 0, 64)
	}
	return
}

// ToFloat64Pure converts any to a float64 value.
//
// Supports the types as follow:
//
//	~bool
//	~string: => strconv.ParseFloat
//	~float32, ~float64
//	~int, ~int8, ~int16, ~int32, ~int64
//	~uint, ~uint8, ~uint16, ~uint32, ~uint64, ~uintptr
//	time.Duration: => F<s>
//
// And the pointer to types above, and the types as follow:
//
//	nil
//	[]byte
//	fmt.Stringer
//	interface{ Float64() float64 }
//	interface{ Float() float64 }
func ToFloat64Pure(any interface{}) (dst float64, err error) {
	switch src := any.(type) {
	case nil:
	case bool:
		if src {
			dst = 1
		}
	case string:
		dst, err = parseFloat64(src)
	case []byte:
		dst, err = parseFloat64(string(src))
	case float32:
		dst = float64(src)
	case float64:
		dst = src
	case int:
		dst = float64(src)
	case int8:
		dst = float64(src)
	case int16:
		dst = float64(src)
	case int32:
		dst = float64(src)
	case int64:
		dst = float64(src)
	case uint:
		dst = float64(src)
	case uint8:
		dst = float64(src)
	case uint16:
		dst = float64(src)
	case uint32:
		dst = float64(src)
	case uint64:
		dst = float64(src)
	case uintptr:
		dst = float64(src)
	case time.Duration:
		dst = float64(src) / float64(time.Second)
	case *time.Duration:
		dst = float64(*src) / float64(time.Second)
	case interface{ Float64() float64 }:
		dst = src.Float64()
	case interface{ Float() float64 }:
		dst = src.Float()
	case fmt.Stringer:
		dst, err = parseFloat64(src.String())
	default:
		dst, err = tryReflectToFloat64(reflect.ValueOf(any))
	}
	return
}

func tryReflectToFloat64(src reflect.Value) (dst float64, err error) {
	switch src.Kind() {
	case reflect.Invalid:
	case reflect.Pointer:
		if !src.IsNil() {
			dst, err = tryReflectToFloat64(src.Elem())
		}

	case reflect.Bool:
		if src.Bool() {
			dst = 1
		}

	case reflect.String:
		dst, err = parseFloat64(src.String())

	case reflect.Float32, reflect.Float64:
		dst = src.Float()

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		dst = float64(src.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		dst = float64(src.Uint())

	default:
		err = fmt.Errorf("cast.ToFloat64: unsupport to convert %T to float64", src.Interface())
	}

	return
}

func parseFloat64(src string) (dst float64, err error) {
	if src != "" {
		dst, err = strconv.ParseFloat(src, 64)
	}
	return
}

// ToDurationPure converts any to a time.Duration value.
//
// Supports the types as follow:
//
//	~string: => N<ms> if integer string, else time.ParseDuration
//	~float32, ~float64: => F<s>
//	~int, ~int8, ~int16, ~int32, ~int64: => N<ms>
//	~uint, ~uint8, ~uint16, ~uint32, ~uint64, ~uintptr: => N<ms>
//	time.Duration
//
// And the pointer to types above, and the types as follow:
//
//	nil
//	[]byte
//	fmt.Stringer
//	interface{ Duration() time.Duration }
func ToDurationPure(any interface{}) (dst time.Duration, err error) {
	switch src := any.(type) {
	case nil:
	case string:
		dst, err = parseDuration(src)
	case []byte:
		dst, err = parseDuration(string(src))
	case float32:
		dst = time.Duration(float64(src) * float64(time.Second))
	case float64:
		dst = time.Duration(src * float64(time.Second))
	case int:
		dst = time.Duration(src) * time.Millisecond
	case int8:
		dst = time.Duration(src) * time.Millisecond
	case int16:
		dst = time.Duration(src) * time.Millisecond
	case int32:
		dst = time.Duration(src) * time.Millisecond
	case int64:
		dst = time.Duration(src) * time.Millisecond
	case uint:
		dst = time.Duration(src) * time.Millisecond
	case uint8:
		dst = time.Duration(src) * time.Millisecond
	case uint16:
		dst = time.Duration(src) * time.Millisecond
	case uint32:
		dst = time.Duration(src) * time.Millisecond
	case uint64:
		dst = time.Duration(src) * time.Millisecond
	case uintptr:
		dst = time.Duration(src) * time.Millisecond
	case time.Duration:
		dst = src
	case *time.Duration:
		dst = *src
	case interface{ Duration() time.Duration }:
		dst = src.Duration()
	case fmt.Stringer:
		dst, err = parseDuration(src.String())
	default:
		dst, err = tryReflectToDuration(reflect.ValueOf(any))
	}
	return
}

func tryReflectToDuration(src reflect.Value) (dst time.Duration, err error) {
	switch src.Kind() {
	case reflect.Invalid:
	case reflect.Pointer:
		if !src.IsNil() {
			dst, err = tryReflectToDuration(src.Elem())
		}

	case reflect.String:
		dst, err = parseDuration(src.String())

	case reflect.Float32, reflect.Float64:
		dst = time.Duration(src.Float()) * time.Millisecond

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		dst = time.Duration(src.Int()) * time.Millisecond

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		dst = time.Duration(src.Uint()) * time.Millisecond

	default:
		err = fmt.Errorf("cast.ToDuration: unsupport to convert %T to time.Duration", src.Interface())
	}

	return
}

func parseDuration(src string) (dst time.Duration, err error) {
	_len := len(src)
	if _len == 0 {
		return
	}

	switch src[_len-1] {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		var i int64
		i, err = strconv.ParseInt(src, 10, 64)
		dst = time.Duration(i) * time.Millisecond
	default:
		dst, err = time.ParseDuration(src)
	}

	return
}

// ToTimeInLocationPure converts any to a time.Time value.
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
//	fmt.Stringer
//	interface{ Time() time.Time }
//
// If loc is nil, use defaults.TimeLocation instead.
// If any is a string-like, use TryParseTime to parse it with layouts.
func ToTimeInLocationPure(any interface{}, loc *time.Location, layouts ...string) (dst time.Time, err error) {
	if loc == nil {
		loc = defaults.TimeLocation.Get()
	}

	switch src := any.(type) {
	case nil:
		dst = dst.In(loc)
	case string:
		dst, err = TryParseTime(src, loc, layouts...)
	case []byte:
		dst, err = TryParseTime(string(src), loc, layouts...)
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
	case fmt.Stringer:
		dst, err = TryParseTime(src.String(), loc, layouts...)
	default:
		dst, err = tryReflectToTimeInLocation(reflect.ValueOf(any), loc, layouts...)
	}

	return
}

func tryReflectToTimeInLocation(src reflect.Value, loc *time.Location,
	layouts ...string) (dst time.Time, err error) {
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

func isIntegerString(s string) bool {
	for i, _len := 0, len(s); i < _len; i++ {
		switch s[i] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '+', '-':
		default:
			return false
		}
	}
	return true
}

// TryParseTime tries to parse the string value with the layouts in turn to time.Time.
//
// If loc is nil, use defaults.TimeLocation instead.
// If layouts is empty, use defaults.TimeFormats instead.
// If value is a integer string, it will be parsee as the unix timestamp.
func TryParseTime(value string, loc *time.Location, layouts ...string) (time.Time, error) {
	if loc == nil {
		loc = defaults.TimeLocation.Get()
	}

	switch value {
	case "", "0000-00-00 00:00:00", "0000-00-00 00:00:00.000", "0000-00-00 00:00:00.000000":
		return time.Time{}.In(loc), nil
	}

	if isIntegerString(value) {
		i, err := strconv.ParseInt(value, 10, 64)
		return time.Unix(i, 0).In(loc), err
	}

	if len(layouts) == 0 {
		layouts = defaults.TimeFormats.Get()
		if len(layouts) == 0 {
			panic("TryParseTime: no time format layouts")
		}
	}

	for _, layout := range layouts {
		if t, err := time.ParseInLocation(layout, value, loc); err == nil {
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("unable to parse time '%s'", value)
}
