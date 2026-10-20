package subject

func ReplaceBits(base, replacement, mask uint8) uint8 { return base&^mask | replacement&mask }
