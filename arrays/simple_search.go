package arrays

func ListSearch(list []int, number int) (resultSearch []int) {
	counter := 0

	var result []int

	for _, list_number := range list {
		if list_number == number {
			result = append(result, number)
			counter++
		}
	}
	return result

}
