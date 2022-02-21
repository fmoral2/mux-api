package api

// buble sort an array of int
func BubbleSort() []int {
	n := []int{3, 2, 4, 3}
	for i := 0; i < len(n); i++ {
		for j := 0; j < len(n)-1; j++ {
			if n[j] > n[j+1] {
				n[j], n[j+1] = n[j+1], n[j]
			}
		}
	}
	return n
}
