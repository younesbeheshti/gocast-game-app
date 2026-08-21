package slice

func DoesExist(list []uint, value uint) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}

	return false
}
