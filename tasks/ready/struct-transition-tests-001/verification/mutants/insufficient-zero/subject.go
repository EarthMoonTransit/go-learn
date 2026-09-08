package subject

type Motor struct{ Speed, Drain, Battery, Distance int }

func Advance(m Motor) Motor {
	if m.Battery >= m.Drain {
		m.Battery -= m.Drain
		m.Distance += m.Speed
	}
	if m.Battery < m.Drain {
		return Motor{}
	}
	return m
}
