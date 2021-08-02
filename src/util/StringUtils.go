package util

func StringContains(slice []string, param string) bool {
	for _, str := range slice {
		if str == param {
			return true
		}
	}
	return false
}
