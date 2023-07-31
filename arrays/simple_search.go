package arrays

func ListSearch(list []int, number int) (result []int) {
	list_length := len(list)
	counter := 0

	result = make([]int, list_length)
	for _, list_number := range list {
		if list_number == number {
			result[counter] = number
			counter++
		}
	}
	return result

}
