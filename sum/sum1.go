package sum

func Sum1(param ...int) int {
	sum := 0
	for _, v := range param {
		sum += v
	}
	return sum
}

//func main() {
//	fmt.Println(Sum1(1, 2, 3, 4))
//}
