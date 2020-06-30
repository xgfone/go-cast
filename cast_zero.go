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

func isZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32,
		reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return math.Float64bits(v.Float()) == 0
	case reflect.Complex64, reflect.Complex128:
		c := v.Complex()
		return math.Float64bits(real(c)) == 0 && math.Float64bits(imag(c)) == 0
	case reflect.Array:
		for i := 0; i < v.Len(); i++ {
			if !isZero(v.Index(i)) {
				return false
			}
		}
		return true
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map,
		reflect.Ptr, reflect.Slice, reflect.UnsafePointer:
		return v.IsNil()
	case reflect.String:
		return v.Len() == 0
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			if !isZero(v.Field(i)) {
				return false
			}
		}
		return true
	default:
		// This should never happens, but will act as a safeguard for
		// later, as a default value doesn't makes sense here.
		panic(&reflect.ValueError{Method: "case.IsZero", Kind: v.Kind()})
	}
}

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
		return isZero(reflect.ValueOf(value))
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
			return isZero(v)
		}
	}
}
