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
)

// ToBool converts any to a bool value.
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
//	interface{ Bool() bool }
//	interface{ IsZero() bool }
func ToBool(any any) (dst bool, err error) {
	switch src := any.(type) {
	case nil:
	case bool:
		dst = src
	case string:
		dst, err = parseBool(src)
	case []byte:
		dst, err = parseBoolByte(src)
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

func parseBoolByte(data []byte) (dst bool, err error) {
	switch len(data) {
	case 0:
	case 1:
		switch data[0] {
		case '\x00':
		case '\x01':
			dst = true
		default:
			err = fmt.Errorf("unsupport to convert []byte to bool")
		}
	default:
		err = fmt.Errorf("unsupport to convert []byte to bool")
	}
	return
}
