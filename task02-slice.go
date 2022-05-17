package homework

func reverse(input []int64) (result []int64) {
	var length int = len(result)
	result = make([]int64, 0, length)

	for i := len(input) - 1; i >= 0; i-- {
		result = append(result, input[i])
	}

	return result
}
