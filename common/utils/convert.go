package utils

func MapKeyToSlice(m map[string]interface{}) []string {
	r := make([]string, 0)
	for k := range m {
		r = append(r, k)
	}

	return r
}
