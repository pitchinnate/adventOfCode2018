package utils

func SearchIntsForInt(ints []int, val int) int {
	for index, intVal := range ints {
		if intVal == val {
			return index
		}
	}
	return -1
}
