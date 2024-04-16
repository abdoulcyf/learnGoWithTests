package iteration

func Repeat(char rune, RepeatCount int) string {
	var repeteadChar string
	for i := 0; i < RepeatCount; i++ {
		if char >= 'A' && char <= 'Z' {
			repeteadChar += string(char + 32)
		} else {
			repeteadChar += string(char)
		}

	}
	return repeteadChar
}
