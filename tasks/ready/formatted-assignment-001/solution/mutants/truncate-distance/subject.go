package subject

import (
	"fmt"
	"math"
)

func Welcome(name string) string { return "Welcome, " + name + "!" }
func FormatAssignment(name string, slot int, direction string, distance float64, peer string) string {
	return fmt.Sprintf("%s\nSlot %03d: %s, %.1f m.\nPeer: %s.", Welcome(name), slot, direction, math.Trunc(distance*10)/10, peer)
}
