// Copyright 2020 xgfone
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
	"math"
	"reflect"
)

// IsZero reports whether the value is ZERO.
func IsZero(value interface{}) bool {
	switch v := value.(type) {
	case bool:
		return !v
	case string:
		return v == ""
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64, uintptr,
		float32, float64:
		return v == 0
	case complex64:
		c := complex128(v)
		return math.Float64bits(real(c)) == 0 && math.Float64bits(imag(c)) == 0
	case complex128:
		return math.Float64bits(real(v)) == 0 && math.Float64bits(imag(v)) == 0
	default:
		switch v := reflect.ValueOf(value); v.Kind() {
		case reflect.Array:
			for i := 0; i < v.Len(); i++ {
				if !IsZero(v.Index(i)) {
					return false
				}
			}
			return true
		case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map,
			reflect.Ptr, reflect.Slice, reflect.UnsafePointer:
			return v.IsNil()
		case reflect.Struct:
			for i := 0; i < v.NumField(); i++ {
				if !IsZero(v.Field(i)) {
					return false
				}
			}
			return true
		default:
			panic(&reflect.ValueError{Method: "cast.IsZero", Kind: v.Kind()})
		}
	}
}

// IsEmpty reports whether the value is empty.
func IsEmpty(value interface{}) bool {
	switch value.(type) {
	case bool, string, float32, float64, complex64, complex128,
		int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64, uintptr:
		return IsZero(value)
	default:
		switch v := reflect.ValueOf(value); v.Kind() {
		case reflect.Array, reflect.Map, reflect.Slice:
			return v.Len() == 0
		default:
			return IsZero(value)
		}
	}
}
