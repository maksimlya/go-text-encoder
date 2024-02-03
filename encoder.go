package encoder

func Encode(message string) []byte {
	return []byte(message)
}
func Decode(bytes[]byte) string {
	return string(bytes[:])
}
