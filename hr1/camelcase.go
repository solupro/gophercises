package main

// Complete the camelcase function below.
func camelcase(s string) int32 {
	var result int32

	if len(s) > 0 {
		result = 1
	}

	list := []byte(s)
	for _, c := range list {
		if c >= 65 && c <= 90 {
			result += 1
		}
	}

	return result
}
