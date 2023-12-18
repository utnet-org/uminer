package utils

func IntInSlice(s int, list []int) bool {
	for _, v := range list {
		if v == s {
			return true
		}
	}

	return false
}
