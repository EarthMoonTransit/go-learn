# Implement stable nested String representations

## Context

Complete the specified operations.

## API

```go
type TemperatureUnit int
const(Celsius TemperatureUnit=0;Fahrenheit TemperatureUnit=1)
type Temperature struct{Degree int;Unit TemperatureUnit}
type SpeedUnit int
const(KmPerHour SpeedUnit=0;MilesPerHour SpeedUnit=1)
type Speed struct{Magnitude int;Unit SpeedUnit}
type Reading struct{Location string;Temperature Temperature;Direction string;Speed Speed;Humidity int}
func (u TemperatureUnit) String() string
func (t Temperature) String() string
func (u SpeedUnit) String() string
func (s Speed) String() string
func (r Reading) String() string
```

## Requirements

- Temperature units format as "°C" or "°F"; speed units as "km/h" or "mph". Only declared units are supplied.
- Temperature and Speed format their integer followed by one space and the formatted unit.
- Reading formats "<Location>: <Temperature>, Wind <Direction> at <Speed>, <Humidity>% Humidity".
- String and fmt.Sprint produce identical results, preserving text, zero and negative integers.
- Implement all five value methods without recursive receiver formatting.

## Constraints

- Use only the Go standard library.
