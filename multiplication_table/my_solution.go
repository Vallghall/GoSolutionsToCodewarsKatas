package multiplication_table

func MultiplicationTable(size int) [][]int {
	if size <= 0 {
		return nil
	}
	result := make([][]int, size, size)
	for i := 0; i < size; i++ {
		result[i] = make([]int, size, size)
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			result[i][j] = (j + 1) * (i + 1)
		}
	}
	return result
}
