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
	"math"
	"time"
)

const (
	J2000         = 2451545.0
	SecondsPerDay = 86400.0
	MinutesPerDay = 1440.0
	HoursPerDay   = 24.0
)

// Transforms a julian day into a solar day.
func ToSolarDay(jd float64) uint {
	z := math.Floor(jd + 0.5)
	f := jd + 0.5 - z

	var A, alpha float64
	if z < 2299161 {
		A = z
	} else {
		alpha = math.Floor((z - 1867216.25) / 36524.25)
		A = z + 1 + alpha - math.Floor(alpha/4.0)
	}

	B := A + 1524.0
	C := math.Floor((B - 122.1) / 365.25)
	D := math.Floor(365.25 * C)
	E := math.Floor((B - D) / 30.6001)

	day := uint(B - D - math.Floor(30.6001*E) + f)
	month := func(E float64) uint {
		if E < 14 {
			return uint(E - 1)
		} else {
			return uint(E - 13)
		}
	}(E)
	year := func(C float64, month uint) uint {
		if month > 2 {
			return uint(C - 4716)
		} else {
			return uint(C - 4715)
		}
	}(C, month)
	k := func(year uint) uint {
		if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
			return 1
		} else {
			return 2
		}
	}(year)

	return uint(float64(275*month/9 - k*uint(float64((month+9)/12)) + day - 30))
}

// Transforms a solar datetime into a julian day.
func ToJulianDay(t time.Time) float64 {
	A := (1461 * (t.Year() + 4800 + (int(t.Month())-14)/12)) / 4
	B := (367 * (int(t.Month()) - 2 - 12*((int(t.Month())-14)/12))) / 12
	C := (3 * ((t.Year() + 4900 + (int(t.Month())-14)/12) / 100)) / 4
	D := t.Day() - 32075
	H := float64(t.Hour()-12) / HoursPerDay
	M := float64(t.Minute()) / MinutesPerDay
	S := float64(t.Second()) / SecondsPerDay

	_, offset := t.Zone()
	Z := float64(offset) / SecondsPerDay

	return float64(A+B-C+D) + H + M + S + Z
}

// Transforms a julian day to century.
func ToJulianCentury(jd float64) float64 {
	return jd * 31557600.0 / 3155695200.0
}
