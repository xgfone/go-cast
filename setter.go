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
	"database/sql"
	"fmt"
	"reflect"
	"time"
)

// Set does the best, using the ToXXX function, to set the value of dst to src.
//
// Support the types as follow:
//
//   - *bool
//   - *int
//   - *int8
//   - *int16
//   - *int32
//   - *int64
//   - *uint
//   - *uint8
//   - *uint16
//   - *uint32
//   - *uint64
//   - *uintptr
//   - *string
//   - *float32
//   - *float64
//   - *time.Time
//   - *time.Duration
//   - reflect.Value
//   - interface sql.Scanner
//   - interface { Set(interface{}) error }
func Set(dst, src interface{}) (err error) {
	switch d := dst.(type) {
	case nil:
		return

	case *bool:
		var v bool
		if v, err = ToBool(src); err == nil {
			*d = v
		}

	case *string:
		var v string
		if v, err = ToString(src); err == nil {
			*d = v
		}

	case *float32:
		var v float64
		if v, err = ToFloat64(src); err == nil {
			*d = float32(v)
		}

	case *float64:
		var v float64
		if v, err = ToFloat64(src); err == nil {
			*d = v
		}

	case *int:
		var v int64
		if v, err = ToInt64(src); err == nil {
			*d = int(v)
		}

	case *int8:
		var v int64
		if v, err = ToInt64(src); err == nil {
			*d = int8(v)
		}

	case *int16:
		var v int64
		if v, err = ToInt64(src); err == nil {
			*d = int16(v)
		}

	case *int32:
		var v int64
		if v, err = ToInt64(src); err == nil {
			*d = int32(v)
		}

	case *int64:
		var v int64
		if v, err = ToInt64(src); err == nil {
			*d = v
		}

	case *uint:
		var v uint64
		if v, err = ToUint64(src); err == nil {
			*d = uint(v)
		}

	case *uint8:
		var v uint64
		if v, err = ToUint64(src); err == nil {
			*d = uint8(v)
		}

	case *uint16:
		var v uint64
		if v, err = ToUint64(src); err == nil {
			*d = uint16(v)
		}

	case *uint32:
		var v uint64
		if v, err = ToUint64(src); err == nil {
			*d = uint32(v)
		}

	case *uint64:
		var v uint64
		if v, err = ToUint64(src); err == nil {
			*d = v
		}

	case *uintptr:
		var v uint64
		if v, err = ToUint64(src); err == nil {
			*d = uintptr(v)
		}

	case *time.Duration:
		var v time.Duration
		if v, err = ToDuration(src); err == nil {
			*d = v
		}

	case *time.Time:
		var v time.Time
		if v, err = ToTime(src); err == nil {
			*d = v
		}

	case reflect.Value:
		err = reflectSet(dst, d, src)

	case interface{ Set(interface{}) error }:
		err = d.Set(src)

	case sql.Scanner:
		err = d.Scan(src)

	default:
		err = reflectSet(dst, reflect.ValueOf(dst), src)
	}

	return
}

// reflectSet is the same as Set, which does the best to set the reflect value dst to src.
func reflectSet(orig interface{}, dst reflect.Value, src interface{}) (err error) {
	if !dst.CanSet() {
		if dst.Kind() == reflect.Pointer {
			elem := dst.Elem()
			if !elem.CanSet() {
				return fmt.Errorf("the dst value %T cannot be set", dst.Interface())
			}
			dst = elem
		}
	}

	switch dst.Kind() {
	case reflect.Bool:
		var v bool
		if v, err = ToBool(src); err == nil {
			dst.SetBool(v)
		}

	case reflect.String:
		var v string
		if v, err = ToString(src); err == nil {
			dst.SetString(v)
		}

	case reflect.Float32, reflect.Float64:
		var v float64
		if v, err = ToFloat64(src); err == nil {
			dst.SetFloat(v)
		}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
		var v int64
		if v, err = ToInt64(src); err == nil {
			dst.SetInt(v)
		}

	case reflect.Int64:
		if _, ok := dst.Interface().(time.Duration); ok {
			v, err := ToDuration(src)
			if err != nil {
				return err
			}
			dst.SetInt(int64(v))
		} else {
			v, err := ToInt64(src)
			if err != nil {
				return err
			}
			dst.SetInt(v)
		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		var v uint64
		if v, err = ToUint64(src); err == nil {
			dst.SetUint(v)
		}

	default:
		iface := dst.Addr().Interface()
		switch d := iface.(type) {
		case *time.Time:
			var v time.Time
			if v, err = ToTime(src); err == nil {
				dst.Set(reflect.ValueOf(v))
			}

		case interface{ Set(interface{}) error }:
			err = d.Set(src)

		case sql.Scanner:
			err = d.Scan(src)

		default:
			err = fmt.Errorf("unsupport to set a value to %T(%v)", orig, orig)
		}
	}

	return
}
