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
	"testing"
	"time"
)

func TestToStringToTime(t *testing.T) {
	t1, err := StringToTime("2020-08-27 23:14:30 +0800 CST")
	if err != nil {
		t.Fatal(err)
	} else if unixtime := t1.Unix(); unixtime != 1598541270 {
		t.Error(unixtime)
	}

	layout := "2006-01-02 15:04:05.999999999 -0700 MST"
	t2, err := StringToTimeInLocation(time.UTC, "2020-08-27 23:14:30 +0800 CST",
		time.ANSIC, layout)
	if err != nil {
		t.Fatal(err)
	} else if unixtime := t2.Unix(); unixtime != 1598541270 {
		t.Error(unixtime)
	}
}

func TestToStringMapInt64(t *testing.T) {
	ms, err := ToStringMapInt64(map[string]interface{}{"a": 123, "b": 456})
	if err != nil {
		t.Error(err)
		return
	}

	for k, v := range ms {
		switch k {
		case "a":
			if v != 123 {
				t.Errorf("a: expected '123', but got '%d'", v)
			}
		case "b":
			if v != 456 {
				t.Errorf("b: expected '456', but got '%d'", v)
			}
		default:
			t.Errorf("error key-value: %s=%d", k, v)
		}
	}

	ms, err = ToStringMapInt64(map[string]int{"a": 123, "b": 456})
	if err != nil {
		t.Error(err)
		return
	}

	for k, v := range ms {
		switch k {
		case "a":
			if v != 123 {
				t.Errorf("a: expected '123', but got '%d'", v)
			}
		case "b":
			if v != 456 {
				t.Errorf("b: expected '456', but got '%d'", v)
			}
		default:
			t.Errorf("error key-value: %s=%d", k, v)
		}
	}
}

func TestIsZero(t *testing.T) {
	if !IsZero(0) || IsZero(1) {
		t.Error()
	}

	if !IsZero(int(0)) || IsZero(int(1)) {
		t.Error()
	}

	if !IsZero(int8(0)) || IsZero(int8(1)) {
		t.Error()
	}

	if !IsZero(int16(0)) || IsZero(int16(1)) {
		t.Error()
	}

	if !IsZero(int32(0)) || IsZero(int32(1)) {
		t.Error()
	}

	if !IsZero(int64(0)) || IsZero(int64(1)) {
		t.Error()
	}

	if !IsZero(uint(0)) || IsZero(uint(1)) {
		t.Error()
	}

	if !IsZero(uint8(0)) || IsZero(uint8(1)) {
		t.Error()
	}

	if !IsZero(uint16(0)) || IsZero(uint16(1)) {
		t.Error()
	}

	if !IsZero(uint32(0)) || IsZero(uint32(1)) {
		t.Error()
	}

	if !IsZero(uint64(0)) || IsZero(uint64(1)) {
		t.Error()
	}

	if !IsZero(float32(0)) || IsZero(float32(1)) {
		t.Error()
	}

	if !IsZero(float64(0)) || IsZero(float64(1)) {
		t.Error()
	}

	if !IsZero(0.0) || IsZero(1.0) {
		t.Error()
	}

	if !IsZero("") || IsZero("1") {
		t.Error()
	}

	if !IsZero(time.Time{}) || IsZero(time.Now()) {
		t.Error()
	}
}

func TestIsSet(t *testing.T) {
	if IsSet(int8(0)) {
		t.Error()
	}
	if !IsSet(int8(1)) {
		t.Error()
	}
}
