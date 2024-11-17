# celestia

A Go library for calculating planetary positions in the Solar system.

## Data

### Planet Enumeration

These enumerations are used throughout the module:

| enum | planet  |
|------|---------|
| 0    | Mercury |
| 1    | Venus   |
| 2    | Earth   |
| 3    | Mars    |
| 4    | Jupiter |
| 5    | Saturn  |

## Functions

### MeanAnomaly (M)

The position that the planet would have relative to its perihelion if the orbit
of the planet were a circle.

| parameter | description |
|-----------|-------------|
| jd        | julian day  |
| p         | planet enum |

### EquationOfCenter (C)

The angular difference between the actual position of a body in its elliptical
orbit and the position it would occupy if its motion were uniform, in a circular
orbit of the same period.

| parameter | description  |
|-----------|--------------|
| jd        | julian day   |
| p         | planet enum  |

### PerihelionLongitude (w)

The sum of the longitude of ascending node (measured on the ecliptic plane) and
the argument of periapsis (measured on the orbital plane).

| parameter | description |
|-----------|-------------|
| p         | planet enum |

### ObliquityEcliptic (e)

The angle between the ecliptic and the celestial equator of the planet.

| parameter | description |
|-----------|-------------|
| p         | planet enum |

### EclipticLongitude (l)

The position along the ecliptic relative to the vernal equinox.

| parameter | description |
|-----------|-------------|
| jd        | julian day  |
| p         | planet enum |

### RightAscension (a)

The right ascension is the coordinate from the equatorial coordinate system in
the sky that has the same role as the longitude in other coordinate systems.

| parameter | description  |
|-----------|--------------|
| jd        | julian day   |
| p         | planet enum  |

### Declination (d)

The declination is the coordinate in the equatorial coordinate system in the sky
that is similar to latitude on Earth. It ranges between −90° at the southern
celestial pole and 90° at the northern celestial pole and is 0° at the celestial
equator.

| parameter | description  |
|-----------|--------------|
| jd        | julian day   |
| p         | planet enum  |

### SiderealTime (theta)

Sidereal time is the rotational angle of the planet at your location, relative
to the stars and the right ascension that is on the celestial meridian at that
moment.

| parameter | description      |
|-----------|------------------|
| jd        | julian day       |
| p         | planet enum      |
| lon       | longitude (west) |

### HourAngle (H)

Hour angle of a celestial body is the difference in right ascension between that
body and the meridian (of right ascension) that is due south at that time,
indicating how long ago (measured in sidereal time) the celestial body passed
through the celestial meridian.

| parameter | description       |
|-----------|-------------------|
| jd        | julian day        |
| p         | planet enum       |
| lon       | longitude (west)  |

### Azimuth (A)

Azimuth is the coordinate from the horizontal coordinate system that indicates
the direction along the horizon. It is convenient to set 0° in the south and to
measure azimuth between −180° and 180°.

| parameter | description       |
|-----------|-------------------|
| jd        | julian day        |
| p         | planet enum       |
| lat       | latitude (north)  |
| lon       | longitude (west)  |

### Altitude (h)

Altitude indicates how high above the horizon a celestial body is. It is 0° at
the horizon, 90° at its zenith (straight up) and -90° in the nadir (straight
down).

| parameter | description       |
|-----------|-------------------|
| jd        | julian day        |
| p         | planet enum       |
| lat       | latitude (north)  |
| lon       | longitude (west)  |

## Sources

- [Astronomy Answers](https://aa.quae.nl/)
- David W. Hughes, B. D. Yallop, C. Y. Hohenkerk, The Equation of Time, Monthly Notices of the Royal Astronomical Society, Volume 238, Issue 4, June 1989, Pages 1529–1535, [https://doi.org/10.1093/mnras/238.4.1529](https://doi.org/10.1093/mnras/238.4.1529)
- [NOAA Solar Calculator](https://gml.noaa.gov/grad/solcalc/)
