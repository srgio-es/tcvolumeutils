package utils

func RemoveEmptyFromSlice(slice []string) []string {
	var result []string

	for _, st := range slice {
		if (len(st) > 0) {
			result = append(result, st)
		}
	}

	return result
}

func RemoveLineEndingsFromSlice(slice []string) []string {
	var result []string

	for _, str := range slice {
		result = append(result, RemoveStringLineEndings(str))
	}

	return result
}