package subject

import (
	"fmt"
)

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

type Temperature struct {
	Degree int
	Unit   TemperatureUnit
}
type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

type Speed struct {
	Magnitude int
	Unit      SpeedUnit
}
type Reading struct {
	Location    string
	Temperature Temperature
	Direction   string
	Speed       Speed
	Humidity    int
}

func (u TemperatureUnit) String() string {
	if u == Celsius {
		return "°C"
	}
	return "°F"
}
func (t Temperature) String() string { return fmt.Sprintf("%d%s", t.Degree, t.Unit) }
func (u SpeedUnit) String() string {
	if u == KmPerHour {
		return "km/h"
	}
	return "mph"
}
func (s Speed) String() string { return fmt.Sprintf("%d%s", s.Magnitude, s.Unit) }
func (r Reading) String() string {
	return fmt.Sprintf("%s: %s, Wind %s at %s, %d%% Humidity", r.Location, r.Temperature, r.Direction, r.Speed, r.Humidity)
}
