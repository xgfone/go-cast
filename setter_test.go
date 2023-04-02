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
	"testing"
	"time"
)

type fmtStringer struct{ v string }

func newStringer(v interface{}) fmtStringer { return fmtStringer{fmt.Sprint(v)} }
func (s fmtStringer) String() string        { return s.v }
func (s *fmtStringer) Set(src interface{}) (err error) {
	s.v, err = ToString(src)
	return
}

func TestSet(t *testing.T) {
	Set(nil, nil)

	var b bool
	testSet(t, &b, "bool1", true, true)
	testSet(t, &b, "bool2", false, false)
	testSet(t, &b, "bool3", "true", true)
	testSet(t, &b, "bool4", []byte("false"), false)
	testSet(t, &b, "bool5", float32(1), true)
	testSet(t, &b, "bool6", float64(0), false)
	testSet(t, &b, "bool7", int(1), true)
	testSet(t, &b, "bool8", int8(0), false)
	testSet(t, &b, "bool9", int16(1), true)
	testSet(t, &b, "bool10", int32(0), false)
	testSet(t, &b, "bool11", int64(1), true)
	testSet(t, &b, "bool12", uint(0), false)
	testSet(t, &b, "bool13", uint8(1), true)
	testSet(t, &b, "bool14", uint16(0), false)
	testSet(t, &b, "bool15", uint32(1), true)
	testSet(t, &b, "bool16", uint64(0), false)
	testSet(t, &b, "bool17", newStringer(1), true)

	var s string
	testSet(t, &s, "string1", true, "true")
	testSet(t, &s, "string2", false, "false")
	testSet(t, &s, "string3", "100", "100")
	testSet(t, &s, "string4", []byte("101"), "101")
	testSet(t, &s, "string5", float32(102), "102")
	testSet(t, &s, "string6", float64(103), "103")
	testSet(t, &s, "string7", int(104), "104")
	testSet(t, &s, "string8", int8(105), "105")
	testSet(t, &s, "string9", int16(106), "106")
	testSet(t, &s, "string10", int32(107), "107")
	testSet(t, &s, "string11", int64(108), "108")
	testSet(t, &s, "string12", uint(109), "109")
	testSet(t, &s, "string13", uint8(110), "110")
	testSet(t, &s, "string14", uint16(111), "111")
	testSet(t, &s, "string15", uint32(112), "112")
	testSet(t, &s, "string16", uint64(113), "113")
	testSet(t, &s, "string17", newStringer(114), "114")

	var f32 float32
	testSet(t, &f32, "float32_1", true, float32(1))
	testSet(t, &f32, "float32_2", false, float32(0))
	testSet(t, &f32, "float32_3", "100", float32(100))
	testSet(t, &f32, "float32_4", []byte("101"), float32(101))
	testSet(t, &f32, "float32_5", float32(102), float32(102))
	testSet(t, &f32, "float32_6", float64(103), float32(103))
	testSet(t, &f32, "float32_7", int(104), float32(104))
	testSet(t, &f32, "float32_8", int8(105), float32(105))
	testSet(t, &f32, "float32_9", int16(106), float32(106))
	testSet(t, &f32, "float32_10", int32(107), float32(107))
	testSet(t, &f32, "float32_11", int64(108), float32(108))
	testSet(t, &f32, "float32_12", uint(109), float32(109))
	testSet(t, &f32, "float32_13", uint8(110), float32(110))
	testSet(t, &f32, "float32_14", uint16(111), float32(111))
	testSet(t, &f32, "float32_15", uint32(112), float32(112))
	testSet(t, &f32, "float32_16", uint64(113), float32(113))
	testSet(t, &f32, "float32_17", newStringer(114), float32(114))

	var f64 float64
	testSet(t, &f64, "float64_1", true, float64(1))
	testSet(t, &f64, "float64_2", false, float64(0))
	testSet(t, &f64, "float64_3", "100", float64(100))
	testSet(t, &f64, "float64_4", []byte("101"), float64(101))
	testSet(t, &f64, "float64_5", float32(102), float64(102))
	testSet(t, &f64, "float64_6", float64(103), float64(103))
	testSet(t, &f64, "float64_7", int(104), float64(104))
	testSet(t, &f64, "float64_8", int8(105), float64(105))
	testSet(t, &f64, "float64_9", int16(106), float64(106))
	testSet(t, &f64, "float64_10", int32(107), float64(107))
	testSet(t, &f64, "float64_11", int64(108), float64(108))
	testSet(t, &f64, "float64_12", uint(109), float64(109))
	testSet(t, &f64, "float64_13", uint8(110), float64(110))
	testSet(t, &f64, "float64_14", uint16(111), float64(111))
	testSet(t, &f64, "float64_15", uint32(112), float64(112))
	testSet(t, &f64, "float64_16", uint64(113), float64(113))
	testSet(t, &f64, "float64_17", newStringer(114), float64(114))
	testSet(t, &f64, "float64_18", time.Second, float64(1))

	var i int
	testSet(t, &i, "int_1", true, int(1))
	testSet(t, &i, "int_2", false, int(0))
	testSet(t, &i, "int_3", "100", int(100))
	testSet(t, &i, "int_4", []byte("101"), int(101))
	testSet(t, &i, "int_5", float32(102), int(102))
	testSet(t, &i, "int_6", float64(103), int(103))
	testSet(t, &i, "int_7", int(104), int(104))
	testSet(t, &i, "int_8", int8(105), int(105))
	testSet(t, &i, "int_9", int16(106), int(106))
	testSet(t, &i, "int_10", int32(107), int(107))
	testSet(t, &i, "int_11", int64(108), int(108))
	testSet(t, &i, "int_12", uint(109), int(109))
	testSet(t, &i, "int_13", uint8(110), int(110))
	testSet(t, &i, "int_14", uint16(111), int(111))
	testSet(t, &i, "int_15", uint32(112), int(112))
	testSet(t, &i, "int_16", uint64(113), int(113))
	testSet(t, &i, "int_17", newStringer(114), int(114))

	var i8 int8
	testSet(t, &i8, "int8_1", true, int8(1))
	testSet(t, &i8, "int8_2", false, int8(0))
	testSet(t, &i8, "int8_3", "100", int8(100))
	testSet(t, &i8, "int8_4", []byte("101"), int8(101))
	testSet(t, &i8, "int8_5", float32(102), int8(102))
	testSet(t, &i8, "int8_6", float64(103), int8(103))
	testSet(t, &i8, "int8_7", int(104), int8(104))
	testSet(t, &i8, "int8_8", int8(105), int8(105))
	testSet(t, &i8, "int8_9", int16(106), int8(106))
	testSet(t, &i8, "int8_10", int32(107), int8(107))
	testSet(t, &i8, "int8_11", int64(108), int8(108))
	testSet(t, &i8, "int8_12", uint(109), int8(109))
	testSet(t, &i8, "int8_13", uint8(110), int8(110))
	testSet(t, &i8, "int8_14", uint16(111), int8(111))
	testSet(t, &i8, "int8_15", uint32(112), int8(112))
	testSet(t, &i8, "int8_16", uint64(113), int8(113))
	testSet(t, &i8, "int8_17", newStringer(114), int8(114))

	var i16 int16
	testSet(t, &i16, "int16_1", true, int16(1))
	testSet(t, &i16, "int16_2", false, int16(0))
	testSet(t, &i16, "int16_3", "100", int16(100))
	testSet(t, &i16, "int16_4", []byte("101"), int16(101))
	testSet(t, &i16, "int16_5", float32(102), int16(102))
	testSet(t, &i16, "int16_6", float64(103), int16(103))
	testSet(t, &i16, "int16_7", int(104), int16(104))
	testSet(t, &i16, "int16_8", int8(105), int16(105))
	testSet(t, &i16, "int16_9", int16(106), int16(106))
	testSet(t, &i16, "int16_10", int32(107), int16(107))
	testSet(t, &i16, "int16_11", int64(108), int16(108))
	testSet(t, &i16, "int16_12", uint(109), int16(109))
	testSet(t, &i16, "int16_13", uint8(110), int16(110))
	testSet(t, &i16, "int16_14", uint16(111), int16(111))
	testSet(t, &i16, "int16_15", uint32(112), int16(112))
	testSet(t, &i16, "int16_16", uint64(113), int16(113))
	testSet(t, &i16, "int16_17", newStringer(114), int16(114))

	var i32 int32
	testSet(t, &i32, "int32_1", true, int32(1))
	testSet(t, &i32, "int32_2", false, int32(0))
	testSet(t, &i32, "int32_3", "100", int32(100))
	testSet(t, &i32, "int32_4", []byte("101"), int32(101))
	testSet(t, &i32, "int32_5", float32(102), int32(102))
	testSet(t, &i32, "int32_6", float64(103), int32(103))
	testSet(t, &i32, "int32_7", int(104), int32(104))
	testSet(t, &i32, "int32_8", int8(105), int32(105))
	testSet(t, &i32, "int32_9", int16(106), int32(106))
	testSet(t, &i32, "int32_10", int32(107), int32(107))
	testSet(t, &i32, "int32_11", int64(108), int32(108))
	testSet(t, &i32, "int32_12", uint(109), int32(109))
	testSet(t, &i32, "int32_13", uint8(110), int32(110))
	testSet(t, &i32, "int32_14", uint16(111), int32(111))
	testSet(t, &i32, "int32_15", uint32(112), int32(112))
	testSet(t, &i32, "int32_16", uint64(113), int32(113))
	testSet(t, &i32, "int32_17", newStringer(114), int32(114))

	var i64 int64
	testSet(t, &i64, "int64_1", true, int64(1))
	testSet(t, &i64, "int64_2", false, int64(0))
	testSet(t, &i64, "int64_3", "100", int64(100))
	testSet(t, &i64, "int64_4", []byte("101"), int64(101))
	testSet(t, &i64, "int64_5", float32(102), int64(102))
	testSet(t, &i64, "int64_6", float64(103), int64(103))
	testSet(t, &i64, "int64_7", int(104), int64(104))
	testSet(t, &i64, "int64_8", int8(105), int64(105))
	testSet(t, &i64, "int64_9", int16(106), int64(106))
	testSet(t, &i64, "int64_10", int32(107), int64(107))
	testSet(t, &i64, "int64_11", int64(108), int64(108))
	testSet(t, &i64, "int64_12", uint(109), int64(109))
	testSet(t, &i64, "int64_13", uint8(110), int64(110))
	testSet(t, &i64, "int64_14", uint16(111), int64(111))
	testSet(t, &i64, "int64_15", uint32(112), int64(112))
	testSet(t, &i64, "int64_16", uint64(113), int64(113))
	testSet(t, &i64, "int64_17", newStringer(114), int64(114))

	var u uint
	testSet(t, &u, "uint_1", true, uint(1))
	testSet(t, &u, "uint_2", false, uint(0))
	testSet(t, &u, "uint_3", "100", uint(100))
	testSet(t, &u, "uint_4", []byte("101"), uint(101))
	testSet(t, &u, "uint_5", float32(102), uint(102))
	testSet(t, &u, "uint_6", float64(103), uint(103))
	testSet(t, &u, "uint_7", int(104), uint(104))
	testSet(t, &u, "uint_8", int8(105), uint(105))
	testSet(t, &u, "uint_9", int16(106), uint(106))
	testSet(t, &u, "uint_10", int32(107), uint(107))
	testSet(t, &u, "uint_11", int64(108), uint(108))
	testSet(t, &u, "uint_12", uint(109), uint(109))
	testSet(t, &u, "uint_13", uint8(110), uint(110))
	testSet(t, &u, "uint_14", uint16(111), uint(111))
	testSet(t, &u, "uint_15", uint32(112), uint(112))
	testSet(t, &u, "uint_16", uint64(113), uint(113))
	testSet(t, &u, "uint_17", newStringer(114), uint(114))

	var u8 uint8
	testSet(t, &u8, "uint8_1", true, uint8(1))
	testSet(t, &u8, "uint8_2", false, uint8(0))
	testSet(t, &u8, "uint8_3", "100", uint8(100))
	testSet(t, &u8, "uint8_4", []byte("101"), uint8(101))
	testSet(t, &u8, "uint8_5", float32(102), uint8(102))
	testSet(t, &u8, "uint8_6", float64(103), uint8(103))
	testSet(t, &u8, "uint8_7", int(104), uint8(104))
	testSet(t, &u8, "uint8_8", int8(105), uint8(105))
	testSet(t, &u8, "uint8_9", int16(106), uint8(106))
	testSet(t, &u8, "uint8_10", int32(107), uint8(107))
	testSet(t, &u8, "uint8_11", int64(108), uint8(108))
	testSet(t, &u8, "uint8_12", uint(109), uint8(109))
	testSet(t, &u8, "uint8_13", uint8(110), uint8(110))
	testSet(t, &u8, "uint8_14", uint16(111), uint8(111))
	testSet(t, &u8, "uint8_15", uint32(112), uint8(112))
	testSet(t, &u8, "uint8_16", uint64(113), uint8(113))
	testSet(t, &u8, "uint8_17", newStringer(114), uint8(114))

	var u16 uint16
	testSet(t, &u16, "uint16_1", true, uint16(1))
	testSet(t, &u16, "uint16_2", false, uint16(0))
	testSet(t, &u16, "uint16_3", "100", uint16(100))
	testSet(t, &u16, "uint16_4", []byte("101"), uint16(101))
	testSet(t, &u16, "uint16_5", float32(102), uint16(102))
	testSet(t, &u16, "uint16_6", float64(103), uint16(103))
	testSet(t, &u16, "uint16_7", int(104), uint16(104))
	testSet(t, &u16, "uint16_8", int8(105), uint16(105))
	testSet(t, &u16, "uint16_9", int16(106), uint16(106))
	testSet(t, &u16, "uint16_10", int32(107), uint16(107))
	testSet(t, &u16, "uint16_11", int64(108), uint16(108))
	testSet(t, &u16, "uint16_12", uint(109), uint16(109))
	testSet(t, &u16, "uint16_13", uint8(110), uint16(110))
	testSet(t, &u16, "uint16_14", uint16(111), uint16(111))
	testSet(t, &u16, "uint16_15", uint32(112), uint16(112))
	testSet(t, &u16, "uint16_16", uint64(113), uint16(113))
	testSet(t, &u16, "uint16_17", newStringer(114), uint16(114))

	var u32 uint32
	testSet(t, &u32, "uint32_1", true, uint32(1))
	testSet(t, &u32, "uint32_2", false, uint32(0))
	testSet(t, &u32, "uint32_3", "100", uint32(100))
	testSet(t, &u32, "uint32_4", []byte("101"), uint32(101))
	testSet(t, &u32, "uint32_5", float32(102), uint32(102))
	testSet(t, &u32, "uint32_6", float64(103), uint32(103))
	testSet(t, &u32, "uint32_7", int(104), uint32(104))
	testSet(t, &u32, "uint32_8", int8(105), uint32(105))
	testSet(t, &u32, "uint32_9", int16(106), uint32(106))
	testSet(t, &u32, "uint32_10", int32(107), uint32(107))
	testSet(t, &u32, "uint32_11", int64(108), uint32(108))
	testSet(t, &u32, "uint32_12", uint(109), uint32(109))
	testSet(t, &u32, "uint32_13", uint8(110), uint32(110))
	testSet(t, &u32, "uint32_14", uint16(111), uint32(111))
	testSet(t, &u32, "uint32_15", uint32(112), uint32(112))
	testSet(t, &u32, "uint32_16", uint64(113), uint32(113))
	testSet(t, &u32, "uint32_17", newStringer(114), uint32(114))

	var u64 uint64
	testSet(t, &u64, "uint64_1", true, uint64(1))
	testSet(t, &u64, "uint64_2", false, uint64(0))
	testSet(t, &u64, "uint64_3", "100", uint64(100))
	testSet(t, &u64, "uint64_4", []byte("101"), uint64(101))
	testSet(t, &u64, "uint64_5", float32(102), uint64(102))
	testSet(t, &u64, "uint64_6", float64(103), uint64(103))
	testSet(t, &u64, "uint64_7", int(104), uint64(104))
	testSet(t, &u64, "uint64_8", int8(105), uint64(105))
	testSet(t, &u64, "uint64_9", int16(106), uint64(106))
	testSet(t, &u64, "uint64_10", int32(107), uint64(107))
	testSet(t, &u64, "uint64_11", int64(108), uint64(108))
	testSet(t, &u64, "uint64_12", uint(109), uint64(109))
	testSet(t, &u64, "uint64_13", uint8(110), uint64(110))
	testSet(t, &u64, "uint64_14", uint16(111), uint64(111))
	testSet(t, &u64, "uint64_15", uint32(112), uint64(112))
	testSet(t, &u64, "uint64_16", uint64(113), uint64(113))
	testSet(t, &u64, "uint64_17", newStringer(114), uint64(114))

	var d time.Duration
	testSet(t, &d, "duration3", "1s", time.Second)
	testSet(t, &d, "duration4", []byte("1m"), time.Minute)
	testSet(t, &d, "duration5", float32(1), time.Second)
	testSet(t, &d, "duration6", float64(2), time.Second*2)
	testSet(t, &d, "duration7", int(104), time.Duration(104)*time.Millisecond)
	testSet(t, &d, "duration8", int8(105), time.Duration(105)*time.Millisecond)
	testSet(t, &d, "duration9", int16(106), time.Duration(106)*time.Millisecond)
	testSet(t, &d, "duration10", int32(107), time.Duration(107)*time.Millisecond)
	testSet(t, &d, "duration11", int64(108), time.Duration(108)*time.Millisecond)
	testSet(t, &d, "duration12", uint(109), time.Duration(109)*time.Millisecond)
	testSet(t, &d, "duration13", uint8(110), time.Duration(110)*time.Millisecond)
	testSet(t, &d, "duration14", uint16(111), time.Duration(111)*time.Millisecond)
	testSet(t, &d, "duration15", uint32(112), time.Duration(112)*time.Millisecond)
	testSet(t, &d, "duration16", uint64(113), time.Duration(113)*time.Millisecond)
	testSet(t, &d, "duration17", newStringer(time.Hour), time.Hour)
	testSet(t, reflect.ValueOf(&d), "duration18", "1s", time.Second)

	var _time time.Time
	testSet(t, &_time, "time1", "2022-07-23T05:56:51Z", time.Unix(1658555811, 0))
	testSet(t, &_time, "time2", []byte("2022-07-23T05:56:52Z"), time.Unix(1658555812, 0))
	testSet(t, &_time, "time3", float32(1658555776), time.Unix(1658555776, 0))
	testSet(t, &_time, "time4", float64(1658555814), time.Unix(1658555814, 0))
	testSet(t, &_time, "time5", int(1658555815), time.Unix(1658555815, 0))
	testSet(t, &_time, "time6", int64(1658555816), time.Unix(1658555816, 0))
	testSet(t, &_time, "time7", uint(1658555817), time.Unix(1658555817, 0))
	testSet(t, &_time, "time8", uint64(1658555818), time.Unix(1658555818, 0))
	testSet(t, &_time, "time9", newStringer("2022-07-23T05:56:59Z"), time.Unix(1658555819, 0))
	testSet(t, reflect.ValueOf(&_time), "tim10", "2022-07-23T05:56:51Z", time.Unix(1658555811, 0))

	var striface fmtStringer
	testSet(t, reflect.ValueOf(&striface), "interface1", 123, newStringer(123))
	testSet(t, reflect.ValueOf(&striface), "interface2", errors.New("test"), newStringer("test"))
}

func testSet(t *testing.T, result interface{}, prefix string, set, expect interface{}) {
	testSetResult(t, prefix, result, Set(result, set), expect)
}

func testSetResult(t *testing.T, prefix string, result interface{}, err error, expect interface{}) {
	if err != nil {
		t.Errorf("%s: %s", prefix, err)
		return
	}

	if rv, ok := result.(reflect.Value); ok {
		result = rv.Elem().Interface()
	} else {
		result = reflect.ValueOf(result).Elem().Interface()
	}

	if rt, ok := result.(time.Time); ok {
		if !rt.Equal(expect.(time.Time)) {
			t.Errorf("%s: expect %v, but got %v", prefix, expect, result)
		}
	} else if !reflect.DeepEqual(result, expect) {
		t.Errorf("%s: expect %v, but got %v", prefix, expect, result)
	}
}
