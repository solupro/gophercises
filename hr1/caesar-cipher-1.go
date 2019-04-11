package main

func isUpper(b byte) bool {
	return b >= 'A' && b <= 'Z'
}

func isAlpha(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}

// Complete the caesarCipher function below.
func caesarCipher(s string, k int32) string {
	result := make([]byte, len(s))

	k %= 26
	offset := byte(k)
	for i, c := range []byte(s) {
		if isAlpha(c) {
			up := isUpper(c)
			c += offset
			if (up && c > 'Z') || c > 'z' {
				c -= byte(26)
			}
		}

		result[i] = c
	}

	return string(result)
}
