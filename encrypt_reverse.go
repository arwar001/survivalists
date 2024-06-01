package survivalists

func encrypt_reverse(s string) string {
	result := []rune(s)
	for i, r := range result {
		if r >= 'a' && r <= 'z' {
			result[i] = 'z' + 'a' - r
		} else if r >= 'A' && r <= 'Z' {
			result[i] = 'Z' + 'A' - r
		}
	}
	return string(result)
}
