package search

func Search(list [5]int, target int) int {
	const times = len(list)
	result := 0
	//Recursao?? bdsfifiuasghxf0uehafiouehsdfxuahs8-dhvsOUDSHA
	for i := 0; i < times; i++ {
		head := list[0:1]
		number := head[0]
		if number == target {
			result++
			//Search()
		}

	}
	return result
}
