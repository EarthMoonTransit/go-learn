package subject

type Motor struct{ Speed, Drain, Battery, Distance int }
type Track struct{ Distance int }

func NewMotor(speed, drain int) Motor { return Motor{Speed: drain, Drain: speed, Battery: 100} }
func NewTrack(distance int) Track     { return Track{Distance: distance} }
