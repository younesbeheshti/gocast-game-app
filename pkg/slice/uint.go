package slice

func DoesExist(list []uint, value uint) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}

	return false
}

func MapFromUint64ToUint(l []uint64) []uint {

	result := make([]uint, len(l))
	for i := range l {
		result[i] = uint(l[i])
	}

	return result
}

func MapFromUintToUint64(l []uint) []uint64 {
	result := make([]uint64, len(l))
	for i := range l {
		result[i] = uint64(l[i])
	}
	return result
}
