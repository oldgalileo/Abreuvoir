package abreuvoir

const tableSeperator rune = '/'

func sanitizeKey(key string) string {
	sanitized := []rune(key)
	if sanitized[0] != tableSeperator {
		sanitized = append([]rune{tableSeperator}, sanitized...)
	}
	if sanitized[len(sanitized)-1] == tableSeperator {
		sanitized = sanitized[:len(sanitized)-1]
	}
	return string(sanitized)
}
