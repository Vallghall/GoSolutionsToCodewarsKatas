package make_the_deadfish_swim

func Parse(data string) []int {
	res := []int{}
	val := 0
	for _, datum := range data {
		switch datum {
		case 'i':
			val += 1
		case 'd':
			val -= 1
		case 's':
			val *= val
		case 'o':
			res = append(res, val)
		default:
			continue
		}
	}
	return res
}
