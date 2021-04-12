package module01

// NumInList will return true if num is in the
// list slice, and false otherwise.
func NumInList(list []int, num int) bool {

	if list == nil {
		return false
	}

	for _, x := range list {
		if num == x {
			return true
		}
	}
	return false
}
