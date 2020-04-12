package slice

/**
 * 对切片进行去重
 */
func Unique(slc []interface{}) []interface{} {
	result := make([]interface{}, 0, len(slc))
	temp := map[interface{}]struct{}{}
	for _, value := range slc {
		if _, ok := temp[value]; !ok {
			temp[value] = struct{}{}
			result = append(result, value)
		}
	}
	return result
}
