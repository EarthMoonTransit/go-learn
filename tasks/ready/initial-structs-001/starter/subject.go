package subject

type Motor struct{ Speed, Drain, Battery, Distance int }
type Track struct{ Distance int }

func NewMotor(speed, drain int) Motor { return Motor{} }
func NewTrack(distance int) Track     { return Track{} }
