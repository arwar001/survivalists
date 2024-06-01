package survivalists

func encrypt_rot13(s string) string {
	var str string
	for _, i := range s {
		if i >= 'a' && i <= 'z' {
			i = i + 13
			for i > 'z' {
				i = i - 26
			}
			str += string(i)
		} else if i >= 'A' && i <= 'Z' {
			i = i + 13
			for i > 'Z' {
				i = i - 26
			}
			str += string(i)
		} else {
			str += string(i)
		}
	}
	return str
}

func decrypt_rot13(s string) string {
	var str string
	for _, i := range s {
		if i >= 'a' && i <= 'z' {
			i = i - 13
			for i > 'z' {
				i = i + 26
			}
			str += string(i)
		} else if i >= 'A' && i <= 'Z' {
			i = i - 13
			for i > 'Z' {
				i = i + 26
			}
			str += string(i)
		} else {
			str += string(i)
		}
	}
	return str
}
