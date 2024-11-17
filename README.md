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

| parameter | description            |
|-----------|------------------------|
| jd        | julian day             |
| p         | planet enum            |
| C         | mean anomaly (degrees) |

### EquationOfCenter (C)

The angular difference between the actual position of a body in its elliptical
orbit and the position it would occupy if its motion were uniform, in a circular
orbit of the same period.

| parameter | description                  |
|-----------|------------------------------|
| jd        | julian day                   |
| p         | planet enum                  |
| C         | equation of center (degrees) |

### PerihelionLongitude (w)

The sum of the longitude of ascending node (measured on the ecliptic plane) and
the argument of periapsis (measured on the orbital plane).

| parameter | description                  |
|-----------|------------------------------|
| p         | planet enum                  |
| w         | equation of center (degrees) |

### ObliquityEcliptic (e)

The angle between the ecliptic and the celestial equator of the planet.

| parameter | description                         |
|-----------|-------------------------------------|
| p         | planet enum                         |
| e         | obliquity of the ecliptic (degrees) |

### EclipticLongitude (l)

The position along the ecliptic relative to the vernal equinox.

| parameter | description                  |
|-----------|------------------------------|
| jd        | julian day                   |
| p         | planet enum                  |
| l         | ecliptic longitude (degrees) |

### RightAscension (a)

Right ascension and declination define the position of a celestial object.

| parameter | description               |
|-----------|---------------------------|
| jd        | julian day                |
| p         | planet enum               |
| a         | right ascension (degrees) |

### Declination (d)

Right ascension and declination define the position of a celestial object.

| parameter | description           |
|-----------|-----------------------|
| jd        | julian day            |
| p         | planet enum           |
| d         | declination (degrees) |
