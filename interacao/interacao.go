package interacao

// functions
func Repeat(param string, times int) string {
	var repetitions string
	for i := 0; i < times; i++ {
		repetitions += param
	}
	return repetitions
}
