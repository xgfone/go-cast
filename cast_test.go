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
)

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
