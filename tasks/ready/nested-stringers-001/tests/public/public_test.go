package subject_test

import (
	"fmt"
	"task.local/nested-stringers-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	for _, c := range []struct {
		s    fmt.Stringer
		want string
	}{{subject.Celsius, "°C"}, {subject.Fahrenheit, "°F"}, {subject.KmPerHour, "km/h"}, {subject.MilesPerHour, "mph"}, {subject.Temperature{Degree: 21, Unit: subject.Celsius}, "21 °C"}, {subject.Temperature{Degree: -17, Unit: subject.Fahrenheit}, "-17 °F"}, {subject.Speed{Magnitude: 18, Unit: subject.KmPerHour}, "18 km/h"}, {subject.Speed{Magnitude: 0, Unit: subject.MilesPerHour}, "0 mph"}} {
		if got := c.s.String(); got != c.want {
			t.Errorf("String got %q want %q", got, c.want)
		}
		if got := fmt.Sprint(c.s); got != c.want {
			t.Errorf("Sprint got %q want %q", got, c.want)
		}
	}
}

func TestPublicCase02(t *testing.T) {
	r := subject.Reading{Location: "Athens", Temperature: subject.Temperature{Degree: 21, Unit: subject.Celsius}, Direction: "N", Speed: subject.Speed{Magnitude: 16, Unit: subject.KmPerHour}, Humidity: 63}
	want := "Athens: 21 °C, Wind N at 16 km/h, 63% Humidity"
	if got := r.String(); got != want {
		t.Errorf("reading got %q want %q", got, want)
	}
	if got := fmt.Sprint(r); got != want {
		t.Errorf("formatted got %q want %q", got, want)
	}
}

func TestPublicCase03(t *testing.T) {
	r := subject.Reading{Location: " x ", Temperature: subject.Temperature{Degree: 0, Unit: subject.Fahrenheit}, Direction: "", Speed: subject.Speed{Magnitude: -2, Unit: subject.MilesPerHour}, Humidity: 0}
	want := " x : 0 °F, Wind  at -2 mph, 0% Humidity"
	if got := fmt.Sprint(r); got != want {
		t.Errorf("boundary got %q want %q", got, want)
	}
}
