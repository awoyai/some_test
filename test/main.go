package main

import "fmt"

func main() {
	a := []int{2, 324, 34, 3125, 67, 4325,1,3,45,6}
	var left, right int = 0, len(a) - 1
	for left < right {
		if a[left] & 1 == 0 {
			swap(&a, left, right)
			right --
			continue
		}
		left ++
	}
	fmt.Println(a)
}

func swap(slice *[]int, a, b int) {
	temp := (*slice)[a]
	(*slice)[a] = (*slice)[b]
	(*slice)[b] = temp 
}
