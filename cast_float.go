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
)

// ToFloat64 converts any to a float64 value.
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
//	interface{ Float64() float64 }
func ToFloat64(any any) (dst float64, err error) {
	switch src := any.(type) {
	case nil:
	case bool:
		if src {
			dst = 1
		}
	case string:
		dst, err = parseFloat64(src)
	case []byte:
		dst, err = parseFloat64Byte(src)
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

func parseFloat64Byte(data []byte) (dst float64, err error) {
	switch len(data) {
	case 0:
	case 1:
		switch data[0] {
		case '\x00':
		case '\x01':
			dst = 1
		default:
			err = fmt.Errorf("unsupport to convert []byte to float64")
		}
	default:
		err = fmt.Errorf("unsupport to convert []byte to float64")
	}
	return
}
