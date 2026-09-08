package subject

type Motor struct{ Speed, Drain, Battery, Distance int }

func Advance(m Motor) Motor {
	if m.Battery >= m.Drain {
		m.Battery -= m.Drain
		m.Distance += m.Speed
	}
	return m
}
