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

// ToString converts any to a string value.
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
func ToString(any any) (dst string, err error) {
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
