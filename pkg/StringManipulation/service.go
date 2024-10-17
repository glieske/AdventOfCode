package stringmanipulation

import "errors"

func IsStringRuneANumber(r rune) bool {
	return r >= '0' && r <= '9'
}

func GetFirstNumberFromString(s string) (int, error) {
	for _, r := range s {
		if IsStringRuneANumber(r) {
			return int(r - '0'), nil
		}
	}
	return -1, errors.New("no number found in string")
}

func GetLastNumberFromString(s string) (int, error) {
	for i := len(s) - 1; i >= 0; i-- {
		if IsStringRuneANumber(rune(s[i])) {
			return int(s[i] - '0'), nil
		}
	}
	return -1, errors.New("no number found in string")
}
