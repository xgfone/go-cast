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
	"strconv"
	"testing"
)

func TestIsIntegerString(t *testing.T) {
	const s = "-1000"
	if !isIntegerString(s) {
		t.Errorf("%s: expect true, but got false", s)
	} else if v, err := strconv.ParseInt(s, 10, 64); err != nil {
		t.Error(err)
	} else if v != -1000 {
		t.Errorf("expect %d, but got %d", -1000, v)
	}

	if isIntegerString("2006-01-02") {
		t.Error("2006-01-02: expect false, but got true")
	}
}
