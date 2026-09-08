package subject

import (
	"fmt"
)

func Welcome(name string) string { return "Welcome, " + name + "!" }
func FormatAssignment(name string, slot int, direction string, distance float64, peer string) string {
	return fmt.Sprintf("%s\nSlot %03d: %s, %.1f m.\nPeer: %s.", Welcome(name), slot, direction, distance, peer)
}
