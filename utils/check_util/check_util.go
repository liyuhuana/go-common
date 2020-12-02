package check_util

func If(condition bool, a, b interface{}) interface{} {
	if condition {
		return a
	}
	return b
}
