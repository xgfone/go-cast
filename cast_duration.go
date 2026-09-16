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
	"time"
)

// ToDuration converts any to a time.Duration value.
//
// Supports the types as follow:
//
//	~string: => time.ParseDuration
//	~float32, ~float64: => F<s>
//	~int, ~int8, ~int16, ~int32, ~int64: => N<ms>
//	~uint, ~uint8, ~uint16, ~uint32, ~uint64, ~uintptr: => N<ms>
//	time.Duration
//
// And the pointer to types above, and the types as follow:
//
//	nil
//	interface{ Duration() time.Duration }
func ToDuration(any any) (dst time.Duration, err error) {
	switch src := any.(type) {
	case nil:
	case string:
		dst, err = parseDuration(src)
	case float32:
		dst = time.Duration(float64(src) * float64(time.Second))
	case float64:
		dst = time.Duration(src * float64(time.Second))
	case int:
		dst = convertIntegerToDuration(src)
	case int8:
		dst = convertIntegerToDuration(src)
	case int16:
		dst = convertIntegerToDuration(src)
	case int32:
		dst = convertIntegerToDuration(src)
	case int64:
		dst = convertIntegerToDuration(src)
	case uint:
		dst = convertIntegerToDuration(src)
	case uint8:
		dst = convertIntegerToDuration(src)
	case uint16:
		dst = convertIntegerToDuration(src)
	case uint32:
		dst = convertIntegerToDuration(src)
	case uint64:
		dst = convertIntegerToDuration(src)
	case uintptr:
		dst = convertIntegerToDuration(src)
	case time.Duration:
		dst = src
	case *time.Duration:
		dst = *src
	case interface{ Duration() time.Duration }:
		dst = src.Duration()
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
		dst = time.Duration(src.Float()) * time.Second

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		dst = convertIntegerToDuration(src.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		dst = convertIntegerToDuration(src.Uint())

	default:
		err = fmt.Errorf("cast.ToDuration: unsupport to convert %T to time.Duration", src.Interface())
	}

	return
}

type _Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

func convertIntegerToDuration[T _Integer](src T) time.Duration {
	return time.Duration(src) * time.Millisecond
}

func parseDuration(src string) (dst time.Duration, err error) {
	if src != "" {
		dst, err = time.ParseDuration(src)
	}
	return
}
