package slice

/**
 * 删除指定位置的元素
 */
func Remove(slc []interface{}, i int) []interface{} {
	// copy([1 2 3], [2 3]) = [2 3 3]
	copy(slc[i:], slc[i+1:])
	return slc[:len(slc)-1]
}
