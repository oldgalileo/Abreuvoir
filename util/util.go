package util

const tableSeperator rune = '/'

// SanitizeKey ensures that the key does not have any trailing '/'s and starts with a '/'
func SanitizeKey(key string) string {
	sanitized := []rune(key)
	if sanitized[0] != tableSeperator {
		sanitized = append([]rune{tableSeperator}, sanitized...)
	}
	if sanitized[len(sanitized)-1] == tableSeperator {
		sanitized = sanitized[:len(sanitized)-1]
	}
	return string(sanitized)
}
