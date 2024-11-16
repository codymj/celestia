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
	"errors"
	"math"

	"github.com/codymj/celestia/julian"
)

const (
	// MeanAnomaly
	M0Mercury = 174.7948
	M1Mercury = 4.09233445
	M0Venus   = 50.4161
	M1Venus   = 1.60213034
	M0Earth   = 357.5291
	M1Earth   = 0.98560028
	M0Mars    = 19.3730
	M1Mars    = 0.52402068
	M0Jupiter = 20.0202
	M1Jupiter = 0.08308529
	M0Saturn  = 317.0207
	M1Saturn  = 0.03344414

	// ObliquityEcliptic
	EMercury = 0.0351
	EVenus   = 2.6376
	EEarth   = 23.4393
	EMars    = 25.1918
	EJupiter = 3.1189
	ESaturn  = 26.7285

	// PerihelionLongitude
	WMercury = 230.3265
	WVenus   = 73.7576
	WEarth   = 102.9373
	WMars    = 71.0041
	WJupiter = 237.1015
	WSaturn  = 99.4587
	WUranus  = 5.4634
	WNeptune = 182.2100
	WPluto   = 184.5484
)

var (
	ErrInvalidEnum = errors.New("invalid planet enum, see README")
)

// MeanAnomaly (M) calculates the position that the planet would have relative
// to its perihelion if the orbit were a circle.
//
// jd: julian day.
//
// p: enum of planet (see README).
func MeanAnomaly(jd float64, p int) (float64, error) {
	var M float64
	var err error

	switch p {
	case 0:
		M = math.Mod(M0Mercury+M1Mercury*(jd-julian.J2000), 360.0)
	case 1:
		M = math.Mod(M0Venus+M1Venus*(jd-julian.J2000), 360.0)
	case 2:
		M = math.Mod(M0Earth+M1Earth*(jd-julian.J2000), 360.0)
	case 3:
		M = math.Mod(M0Mars+M1Mars*(jd-julian.J2000), 360.0)
	case 4:
		M = math.Mod(M0Jupiter+M1Jupiter*(jd-julian.J2000), 360.0)
	case 5:
		M = math.Mod(M0Saturn+M1Saturn*(jd-julian.J2000), 360.0)
	default:
		err = ErrInvalidEnum
	}

	return M, err
}

// ObliquityEcliptic (e) is the angle between the ecliptic and the celestial
// equator of the planet.
//
// p: enum of planet (see README).
func ObliquityEcliptic(p int) (float64, error) {
	var e float64
	var err error

	switch p {
	case 0:
		e = EMercury
	case 1:
		e = EVenus
	case 2:
		e = EEarth
	case 3:
		e = EMars
	case 4:
		e = EJupiter
	case 5:
		e = ESaturn
	default:
		err = ErrInvalidEnum
	}

	return e, err
}

// PerihelionLongitude (w) is the sum of the longitude of ascending node
// (measured on the ecliptic plane) and the argument of periapsis (measured on
// the orbital plane).
//
//	p: enum of the planet.
func PerihelionLongitude(p int) (float64, error) {
	var w float64
	var err error

	switch p {
	case 0:
		w = WMercury
	case 1:
		w = WVenus
	case 2:
		w = WEarth
	case 3:
		w = WMars
	case 4:
		w = WJupiter
	case 5:
		w = WSaturn
	default:
		err = ErrInvalidEnum
	}

	return w, err
}
