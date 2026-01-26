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

// ToInt64 converts any to an int64 value.
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
//	interface{ Int64() int64 }
func ToInt64(any any) (dst int64, err error) {
	switch src := any.(type) {
	case nil:
	case bool:
		if src {
			dst = 1
		}
	case string:
		dst, err = parseInt64(src)
	case []byte:
		dst, err = parseInt64Byte(src)
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

func parseInt64Byte(data []byte) (dst int64, err error) {
	switch len(data) {
	case 0:
	case 1:
		switch data[0] {
		case '\x00':
		case '\x01':
			dst = 1
		default:
			err = fmt.Errorf("unsupport to convert []byte to int64")
		}
	default:
		err = fmt.Errorf("unsupport to convert []byte to int64")
	}
	return
}
