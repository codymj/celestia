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

package solarposition

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
		{"ForEarth", 2453097.0, 3, 112.65309536000007, nil},
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
		{"ForMars", 3, EMars, nil},
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
		P    float64
		err  error
	}{
		{"ForEarth", 2, PEarth, nil},
		{"ForMars", 3, PMars, nil},
		{"InvalidPlanet", 12, 0, ErrInvalidEnum},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			P, err := PerihelionLongitude(tt.p)
			assert.Equal(t, tt.P, P)
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
		{"ForMars", 2453097.0, 3, 9.409206613394835, nil},
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
		{"ForMars", 2453097.0, 3, 122.0623019733949, nil},
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
		{"ForMars", 2453097.0, 3, 13.066401973394875, nil},
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

// Declination tests.
func TestDeclination(t *testing.T) {
	tests := []struct {
		name string
		jd   float64
		p    int
		d    float64
		err  error
	}{
		{"ForEarth", 2453097.0, 2, 4.740184662324431, nil},
		{"ForMars", 2453097.0, 3, 5.496702418591823, nil},
		{"InvalidPlanet", 2453097.0, 12, 0, ErrInvalidEnum},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d, err := Declination(tt.jd, tt.p)
			assert.Equal(t, tt.d, d)
			assert.Equal(t, tt.err, err)
		})
	}
}

// SiderealTime tests.
func TestSiderealTime(t *testing.T) {
	tests := []struct {
		name  string
		jd    float64
		p     int
		lon   float64
		theta float64
		err   error
	}{
		{"ForEarth", 2453097.0, 2, -5.0, 14.834671999909915, nil},
		{"ForMars", 2453097.0, 3, 184.6, 33.13916751998477, nil},
		{"InvalidPlanet", 2453097.0, 12, 34.7, 0, ErrInvalidEnum},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			theta, err := SiderealTime(tt.jd, tt.p, tt.lon)
			assert.Equal(t, tt.theta, theta)
			assert.Equal(t, tt.err, err)
		})
	}
}

// HourAngle tests.
func TestHourAngle(t *testing.T) {
	tests := []struct {
		name string
		jd   float64
		p    int
		lon  float64
		H    float64
		err  error
	}{
		{"ForEarth", 2453097.0, 2, -5.0, 3.76980128420956, nil},
		{"ForMars", 2453097.0, 3, 184.6, 21.278579105151533, nil},
		{"InvalidPlanet", 2453097.0, 12, 34.7, 0, ErrInvalidEnum},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			H, err := HourAngle(tt.jd, tt.p, tt.lon)
			assert.Equal(t, tt.H, H)
			assert.Equal(t, tt.err, err)
		})
	}
}

// Azimuth tests.
func TestAzimuth(t *testing.T) {
	tests := []struct {
		name string
		jd   float64
		p    int
		lat  float64
		lon  float64
		A    float64
		err  error
	}{
		{"ForEarth", 2453097.0, 2, 52.0, -5.0, 5.109917114922145, nil},
		{"ForMars", 2453097.0, 3, -14.6, 184.6, 132.10875118644827, nil},
		{"InvalidPlanet", 2453097.0, 12, -45, 34.7, 0, ErrInvalidEnum},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			A, err := Azimuth(tt.jd, tt.p, tt.lat, tt.lon)
			assert.Equal(t, tt.A, A)
			assert.Equal(t, tt.err, err)
		})
	}
}

// Azimuth tests.
func TestAltitude(t *testing.T) {
	tests := []struct {
		name string
		jd   float64
		p    int
		lat  float64
		lon  float64
		h    float64
		err  error
	}{
		{"ForEarth", 2453097.0, 2, 52.0, -5.0, 42.63670285961314, nil},
		{"ForMars", 2453097.0, 3, -14.6, 184.6, 60.861557344577825, nil},
		{"InvalidPlanet", 2453097.0, 12, -45, 34.7, 0, ErrInvalidEnum},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h, err := Altitude(tt.jd, tt.p, tt.lat, tt.lon)
			assert.Equal(t, tt.h, h)
			assert.Equal(t, tt.err, err)
		})
	}
}

// Transit tests.
func TestTransit(t *testing.T) {
	tests := []struct {
		name      string
		jd        float64
		p         int
		lon       float64
		J_transit float64
		err       error
	}{
		{"ForEarth", 2453097.0, 2, -5.0, 2.4530969895304884e+06, nil},
		{"ForMars", 2453097.0, 3, 184.6, 2.453096939282806e+06, nil},
		{"InvalidPlanet", 2453097.0, 12, -45, 0, ErrInvalidEnum},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			J_transit, err := TransitTime(tt.jd, tt.p, tt.lon)
			assert.Equal(t, tt.J_transit, J_transit)
			assert.Equal(t, tt.err, err)
		})
	}
}

// Sunrise tests.
func TestSunrise(t *testing.T) {
	tests := []struct {
		name   string
		jd     float64
		p      int
		lat    float64
		lon    float64
		J_rise float64
		err    error
	}{
		{"ForEarth", 2453097.0, 2, 52, -5.0, 2.4530967190208086e+06, nil},
		{"ForMars", 2453097.0, 3, -14.6, 184.6, 2.453096686034785e+06, nil},
		{"InvalidPlanet", 2453097.0, 23, 12, -45, 0, ErrInvalidEnum},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			J_rise, err := SunriseTime(tt.jd, tt.p, tt.lat, tt.lon)
			assert.Equal(t, tt.J_rise, J_rise)
			assert.Equal(t, tt.err, err)
		})
	}
}

// Sunset tests.
func TestSunset(t *testing.T) {
	tests := []struct {
		name  string
		jd    float64
		p     int
		lat   float64
		lon   float64
		J_set float64
		err   error
	}{
		{"ForEarth", 2453097.0, 2, 52, -5.0, 2.4530972600402692e+06, nil},
		{"ForMars", 2453097.0, 3, -14.6, 184.6, 2.453097192530769e+06, nil},
		{"InvalidPlanet", 2453097.0, 23, 12, -45, 0, ErrInvalidEnum},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			J_set, err := SunsetTime(tt.jd, tt.p, tt.lat, tt.lon)
			assert.Equal(t, tt.J_set, J_set)
			assert.Equal(t, tt.err, err)
		})
	}
}
