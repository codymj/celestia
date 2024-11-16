// Copyright 2024 Cody Johnson
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

package julian

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// ToSolarDay test.
func TestToSolarDay(t *testing.T) {
	tests := []struct {
		name string
		jd   float64
		d    uint
	}{
		{"Test1", 2451545.2577546295, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := ToSolarDay(tt.jd)
			assert.Equal(t, tt.d, d)
		})
	}
}

// ToJulian test.
func TestToJulianDay(t *testing.T) {
	tests := []struct {
		name string
		t    string
		jd   float64
	}{
		{"Test1", "2000-01-01T18:11:10-01:00", 2.451545216087963e+06},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts, _ := time.Parse(time.RFC3339, tt.t)
			jd := ToJulianDay(ts)
			assert.Equal(t, tt.jd, jd)
		})
	}
}

// TestToJulianCentury tests ToJulianCentury()
func TestToJulianCentury(t *testing.T) {
	tests := []struct {
		name string
		t    string
		jc   float64
	}{
		{"Test1", "2000-01-01T18:11:10Z", 24515.955985266733},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tm, _ := time.Parse(time.RFC3339, tt.t)
			jd := ToJulianDay(tm)
			jc := ToJulianCentury(jd)
			assert.Equal(t, tt.jc, jc)
		})
	}
}
