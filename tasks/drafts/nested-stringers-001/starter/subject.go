package subject

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

func (u TemperatureUnit) String() string { return "" }
func (t Temperature) String() string     { return "" }
func (u SpeedUnit) String() string       { return "" }
func (s Speed) String() string           { return "" }
func (r Reading) String() string         { return "" }
