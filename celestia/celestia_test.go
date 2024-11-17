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

package celestia

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// MeanAnomaly tests.
func TestMeanAnomaly(t *testing.T) {
	tests := []struct {
		name string
		jd   float64
		p    int
		M    float64
		err  error
	}{
		{"ForEarth", 2453097.0, 2, 87.18073456000002, nil},
		{"InvalidPlanet", 2453097.0, 12, 0, ErrInvalidEnum},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			M, err := MeanAnomaly(tt.jd, tt.p)
			assert.Equal(t, tt.M, M)
			assert.Equal(t, tt.err, err)
		})
	}
}

// ObliquityEcliptic tests.
func TestObliquityEcliptic(t *testing.T) {
	tests := []struct {
		name string
		p    int
		e    float64
		err  error
	}{
		{"ForEarth", 2, EEarth, nil},
		{"InvalidPlanet", 12, 0, ErrInvalidEnum},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e, err := ObliquityEcliptic(tt.p)
			assert.Equal(t, tt.e, e)
			assert.Equal(t, tt.err, err)
		})
	}
}

// PerihelionLongitude tests.
func TestPerihelionLongitude(t *testing.T) {
	tests := []struct {
		name string
		p    int
		w    float64
		err  error
	}{
		{"ForEarth", 2, WEarth, nil},
		{"InvalidPlanet", 12, 0, ErrInvalidEnum},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w, err := PerihelionLongitude(tt.p)
			assert.Equal(t, tt.w, w)
			assert.Equal(t, tt.err, err)
		})
	}
}

// EquationOfCenter tests.
func TestEquationOfCenter(t *testing.T) {
	tests := []struct {
		name string
		jd   float64
		p    int
		C    float64
		err  error
	}{
		{"ForEarth", 2453097.0, 2, 1.9141507379386618, nil},
		{"InvalidPlanet", 2453097.0, 12, 0, ErrInvalidEnum},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			C, err := EquationOfCenter(tt.jd, tt.p)
			assert.Equal(t, tt.C, C)
			assert.Equal(t, tt.err, err)
		})
	}
}

// TrueAnomaly tests.
func TestTrueAnomaly(t *testing.T) {
	tests := []struct {
		name string
		jd   float64
		p    int
		v    float64
		err  error
	}{
		{"ForEarth", 2453097.0, 2, 89.09488529793867, nil},
		{"InvalidPlanet", 2453097.0, 12, 0, ErrInvalidEnum},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, err := TrueAnomaly(tt.jd, tt.p)
			assert.Equal(t, tt.v, v)
			assert.Equal(t, tt.err, err)
		})
	}
}

// EclipticLongitude tests.
func TestEclipticLongitude(t *testing.T) {
	tests := []struct {
		name string
		jd   float64
		p    int
		l    float64
		err  error
	}{
		{"ForEarth", 2453097.0, 2, 12.032185297938668, nil},
		{"InvalidPlanet", 2453097.0, 12, 0, ErrInvalidEnum},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l, err := EclipticLongitude(tt.jd, tt.p)
			assert.Equal(t, tt.l, l)
			assert.Equal(t, tt.err, err)
		})
	}
}

// RightAscension tests.
func TestRightAscension(t *testing.T) {
	tests := []struct {
		name string
		jd   float64
		p    int
		a    float64
		err  error
	}{
		{"ForEarth", 2453097.0, 2, 11.064870715700355, nil},
		{"ForMars", 2453097.0, 3, 11.860588414833234, nil},
		{"InvalidPlanet", 2453097.0, 12, 0, ErrInvalidEnum},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a, err := RightAscension(tt.jd, tt.p)
			assert.Equal(t, tt.a, a)
			assert.Equal(t, tt.err, err)
		})
	}
}
