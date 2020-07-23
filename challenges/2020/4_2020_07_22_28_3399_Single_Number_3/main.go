package main

import "fmt"

// Naive implementation :(
func singleNumber(nums []int) (res []int) {
	unique := map[int]int{}
	for _, n := range nums {
		exists := unique[n]
		if exists == 0 {
			unique[n] = 1
		} else {
			unique[n] = 0
		}
	}
	for k, v := range unique {
		if v == 1 {
			res = append(res, k)
		}
	}
	return
}

func main() {
	fmt.Println(singleNumber([]int{1, 2, 1, 3, 2, 5, 5, 2}))
}
