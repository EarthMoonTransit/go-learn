package definedtypes

type Samples int
type Bytes int

func EncodedSize(count Samples, bytesPerSample int) Bytes {
	return Bytes(count) * Bytes(bytesPerSample)
}

func Combine(left, right Samples) Samples {
	return left + right
}
