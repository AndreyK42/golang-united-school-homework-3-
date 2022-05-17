package homework

func average(input [15]float32) (result float32) {
	var n int = len(input)
	var sum float32 = 0.0

	for i := 0; i < n; i++ {
		sum += input[i]
	}

	return sum / float32(n)
}
