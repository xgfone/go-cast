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
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

// ToUint64 converts any to a uint64 value.
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
//	interface{ Uint64() uint64 }
func ToUint64(any any) (dst uint64, err error) {
	switch src := any.(type) {
	case nil:
	case bool:
		if src {
			dst = 1
		}
	case string:
		dst, err = parseUint64(src)
	case []byte:
		dst, err = parseUint64Byte(src)
	case float32:
		dst, err = convertUint64(src)
	case float64:
		dst, err = convertUint64(src)
	case int:
		dst, err = convertUint64(src)
	case int8:
		dst, err = convertUint64(src)
	case int16:
		dst, err = convertUint64(src)
	case int32:
		dst, err = convertUint64(src)
	case int64:
		dst, err = convertUint64(src)
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

func parseUint64Byte(data []byte) (dst uint64, err error) {
	switch len(data) {
	case 0:
	case 1:
		switch data[0] {
		case '\x00':
		case '\x01':
			dst = 1
		default:
			err = fmt.Errorf("unsupport to convert []byte to uint64")
		}
	default:
		err = fmt.Errorf("unsupport to convert []byte to uint64")
	}
	return
}

type _Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}

var errNegativeNumber = errors.New("cannot convert a negative number to uint64")

func convertUint64[T _Number](src T) (uint64, error) {
	if src < 0 {
		return 0, errNegativeNumber
	}
	return uint64(src), nil
}
