package generator

func Newintgenerator() func() int {
	var total int
	return func() int {
		total++
		return total
	}
}
