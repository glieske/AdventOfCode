package typeconvert

func StringRuneToInt(r rune) int {
	return int(r - '0')
}

func StringRuneToString(r rune) string {
	return string(r)
}

func StringRuneToInt32(r rune) int32 {
	return int32(r - '0')
}

func StringRuneToInt64(r rune) int64 {
	return int64(r - '0')
}
