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
	RAD = math.Pi / 180
	DEG = 180 / math.Pi

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

	// EquationOfCenter
	C1Mercury = 23.4400
	C2Mercury = 2.9818
	C3Mercury = 0.5255
	C4Mercury = 0.1058
	C5Mercury = 0.0241
	C6Mercury = 0.0055
	C1Venus   = 0.7758
	C2Venus   = 0.0033
	C3Venus   = 0.0000
	C4Venus   = 0.0000
	C5Venus   = 0.0000
	C6Venus   = 0.0000
	C1Earth   = 1.9148
	C2Earth   = 0.0200
	C3Earth   = 0.0003
	C4Earth   = 0.0000
	C5Earth   = 0.0000
	C6Earth   = 0.0000
	C1Mars    = 10.6912
	C2Mars    = 0.6228
	C3Mars    = 0.0503
	C4Mars    = 0.0046
	C5Mars    = 0.0005
	C6Mars    = 0.0000
	C1Jupiter = 5.5549
	C2Jupiter = 0.1683
	C3Jupiter = 0.0071
	C4Jupiter = 0.0003
	C5Jupiter = 0.0000
	C6Jupiter = 0.0000
	C1Saturn  = 6.3585
	C2Saturn  = 0.2204
	C3Saturn  = 0.0106
	C4Saturn  = 0.0006
	C5Saturn  = 0.0000
	C6Saturn  = 0.0000

	// SiderealTime
	T0Mercury = 132.3282
	T1Mercury = 6.1385025
	T0Venus   = 104.9067
	T1Venus   = -1.4813688
	T0Earth   = 280.1470
	T1Earth   = 360.9856235
	T0Mars    = 313.3827
	T1Mars    = 350.89198226
	T0Jupiter = 145.9722
	T1Jupiter = 870.5360000
	T0Saturn  = 174.3508
	T1Saturn  = 810.7939024
)

var (
	ErrInvalidEnum = errors.New("invalid planet enum, see README")
)

// Mean anomaly (M) calculates the position that the planet would have relative
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

// Obliquity ecliptic (e) is the angle between the ecliptic and the celestial
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

// Perihelion longitude (w) is the sum of the longitude of ascending node
// (measured on the ecliptic plane) and the argument of periapsis (measured on
// the orbital plane).
//
// p: enum of planet (see README).
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

// Equation of center (C) is the angular difference between the actual position
// of a body in its elliptical orbit and the position it would occupy if its
// motion were uniform.
//
// jd: julian day.
//
// p: enum of planet (see README).
func EquationOfCenter(jd float64, p int) (float64, error) {
	M, err := MeanAnomaly(jd, p)
	if err != nil {
		return 0, err
	}

	var C float64

	calc := func(c1, c2, c3, c4, c5, c6, m float64) float64 {
		return c1*math.Sin(m) + c2*math.Sin(2*m) + c3*math.Sin(3*m) +
			c4*math.Sin(4*m) + c5*math.Sin(5*m) + c6*math.Sin(6*m)
	}

	switch p {
	case 0:
		C = calc(
			C1Mercury, C2Mercury, C3Mercury,
			C4Mercury, C5Mercury, C6Mercury,
			M*RAD,
		)
	case 1:
		C = calc(
			C1Venus, C2Venus, C3Venus,
			C4Venus, C5Venus, C6Venus,
			M*RAD,
		)
	case 2:
		C = calc(
			C1Earth, C2Earth, C3Earth,
			C4Earth, C5Earth, C6Earth,
			M*RAD,
		)
	case 3:
		C = calc(
			C1Mars, C2Mars, C3Mars,
			C4Mars, C5Mars, C6Mars,
			M*RAD,
		)
	case 4:
		C = calc(
			C1Jupiter, C2Jupiter, C3Jupiter,
			C4Jupiter, C5Jupiter, C6Jupiter,
			M*RAD,
		)
	case 5:
		C = calc(
			C1Saturn, C2Saturn, C3Saturn,
			C4Saturn, C5Saturn, C6Saturn,
			M*RAD,
		)
	default:
		err = ErrInvalidEnum
	}

	return C, err
}

// True anomaly (v) is the sum of the mean anomaly (M) and the equation of
// center (C).
//
// jd: julian day.
//
// p: enum of planet (see README).
func TrueAnomaly(jd float64, p int) (float64, error) {
	M, err := MeanAnomaly(jd, p)
	if err != nil {
		return 0, err
	}

	C, err := EquationOfCenter(jd, p)
	if err != nil {
		return 0, err
	}

	return M + C, err
}

// Ecliptic longitude (l) is the position along the ecliptic relative to the
// vernal equinox (in degrees).
//
// jd: julian day.
//
// p: enum of planet (see README).
func EclipticLongitude(jd float64, p int) (float64, error) {
	M, err := MeanAnomaly(jd, p)
	if err != nil {
		return 0, err
	}

	w, err := PerihelionLongitude(p)
	if err != nil {
		return 0, err
	}

	C, err := EquationOfCenter(jd, p)
	if err != nil {
		return 0, err
	}

	L := M + w + 180

	l := L + C
	for l > 360.0 {
		l = math.Mod(l, 360.0)
	}

	return l, err
}

// Right ascension (a) is the angular distance of a celestial object's hour
// circle east of the vernal equinox, measured along the celestial equator (in
// degrees).
//
// jd: julian day.
//
// p: enum of planet (see README).
func RightAscension(jd float64, p int) (float64, error) {
	l, err := EclipticLongitude(jd, p)
	if err != nil {
		return 0, err
	}

	e, err := ObliquityEcliptic(p)
	if err != nil {
		return 0, err
	}

	a := math.Atan2(math.Sin(l*RAD)*math.Cos(e*RAD), math.Cos(l*RAD)) * DEG

	return a, err
}

// Declination (d) determines from which parts of the planet the object can be
// visible.
//
// jd: julian day.
//
// p: enum of the planet (see README).
func Declination(jd float64, p int) (float64, error) {
	l, err := EclipticLongitude(jd, p)
	if err != nil {
		return 0, err
	}

	e, err := ObliquityEcliptic(p)
	if err != nil {
		return 0, err
	}

	d := math.Atan(math.Sin(l*RAD)*math.Sin(e*RAD)) * DEG

	return d, err
}

// Sidereal time (theta) is the rotational angle of the planet at your location,
// relative to the stars and the right ascension that is on the celestial
// meridian at that moment.
//
// jd: julian day.
//
// p: enum of the planet (see README).
//
// lon: longitude (west).
func SiderealTime(jd float64, p int, lon float64) (float64, error) {
	var theta float64
	var err error

	switch p {
	case 0:
		theta = T0Mercury + T1Mercury*(jd-julian.J2000) - lon
	case 1:
		theta = T0Venus + T1Venus*(jd-julian.J2000) - lon
	case 2:
		theta = T0Earth + T1Earth*(jd-julian.J2000) - lon
	case 3:
		theta = T0Mars + T1Mars*(jd-julian.J2000) - lon
	case 4:
		theta = T0Jupiter + T1Jupiter*(jd-julian.J2000) - lon
	case 5:
		theta = T0Saturn + T1Saturn*(jd-julian.J2000) - lon
	default:
		err = ErrInvalidEnum
	}

	for theta > 360.0 {
		theta = math.Mod(theta, 360.0)
	}

	return theta, err
}

// Hour angle (H) of a celestial body is the difference in right ascension
// between that body and the meridian (of right ascension) that is due south at
// that time, indicating how long ago (measured in sidereal time) the celestial
// body passed through the celestial meridian.
//
// jd: julian day.
//
// p: enum of the planet (see README).
//
// lon: longitude (west).
func HourAngle(jd float64, p int, lon float64) (float64, error) {
	theta, err := SiderealTime(jd, p, lon)
	if err != nil {
		return 0, err
	}

	a, err := RightAscension(jd, p)
	if err != nil {
		return 0, err
	}

	return theta - a, err
}

// Azimuth (A) is the coordinate from the horizontal coordinate system that
// indicates the direction along the horizon. It is convenient to set 0° in the
// south and to measure azimuth between −180° and 180°.
//
// jd: julian day.
//
// p: enum of the planet (see README).
//
// lat: latitude (north)
//
// lon: longitude (west).
func Azimuth(jd float64, p int, lat, lon float64) (float64, error) {
	d, err := Declination(jd, p)
	if err != nil {
		return 0, err
	}

	H, err := HourAngle(jd, p, lon)
	if err != nil {
		return 0, err
	}

	A := math.Atan2(
		math.Sin(H*RAD),
		math.Cos(H*RAD)*math.Sin(lat*RAD)-math.Tan(d*RAD)*math.Cos(lat*RAD),
	) * DEG

	return A, err
}

// Altitude (h) indicates how high above the horizon a celestial body is. It is
// 0° at the horizon, 90° at its zenith (straight up) and -90° in the nadir (
// straight down).
//
// jd: julian day.
//
// p: enum of the planet (see README).
//
// lat: latitude (north)
//
// lon: longitude (west).
func Altitude(jd float64, p int, lat, lon float64) (float64, error) {
	d, err := Declination(jd, p)
	if err != nil {
		return 0, err
	}

	H, err := HourAngle(jd, p, lon)
	if err != nil {
		return 0, err
	}

	h := math.Asin(
		math.Sin(lat*RAD)*math.Sin(d*RAD)+
			math.Cos(lat*RAD)*math.Cos(d*RAD)*math.Cos(H*RAD),
	) * DEG

	return h, err
}
