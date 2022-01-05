package compare

//IsVowelOrY checks if is a vowel (inclues y)
func IsVowelOrY(c rune) bool {
	switch c {
	case 'A', 'E', 'I', 'O', 'U', 'Y':
		return true
	default:
		return false
	}
}

//IsVowel checks if is a vowel
func IsVowel(c rune) bool {
	switch c {
	case 'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}
}
