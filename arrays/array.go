package arrays

func ListCounter(list []int) int {
	counter := 0

	for _, number := range list {
		counter += number
	}
	return counter
}
