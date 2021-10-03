package md

func Bytes2(a, b int, e byte) [][]byte {
	s := make([][]byte, a)
	for i := range s {
		s[i] = make([]byte, b)
		for j := range s[i] {
			s[i][j] = e
		}
	}
	return s
}

func Ints2(a, b int, e int) [][]int {
	s := make([][]int, a)
	for i := range s {
		s[i] = make([]int, b)
		for j := range s[i] {
			s[i][j] = e
		}
	}
	return s
}
