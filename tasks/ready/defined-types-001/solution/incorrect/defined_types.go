package definedtypes

type Samples = int
type Bytes = int

func EncodedSize(count Samples, bytesPerSample int) Bytes {
	return count * bytesPerSample
}

func Combine(left, right Samples) Samples {
	return left + right
}
