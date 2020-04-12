package slice

/**
 * 倒序
 */
func Reverse(s []interface{}) {
	for i, j := 0, len(s)-1; i < j; i++ {
		j = len(s) - 1 - i
		s[i], s[j] = s[j], s[i]
	}
}
