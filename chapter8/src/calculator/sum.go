package calculator

// Sum realiza a toma de uma lista de valores inteiros
func Sum(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}
