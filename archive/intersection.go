package main

import "fmt"

// Input: A = [[0,2],[5,10],[|13,23],[24,25]],
// Input  B = [[1,5],[8|,12],[15,24],[25,26]]
// Output: [[1,2],[5,5],[8,10],[15,23],[24,24],[25,25]]
//[[4,6],[7,8],[10,17]]
// [[5,10]]
//[[5,6],[7,8],[10,10]]

// A [[0,4],[7,8],[12,19]]
// B [[0,10],[14,15],[18,20]]
// OUTPUT [[0,4],[7,8],[14,15],[18,19]]

// Reminder: The inputs and the desired output are lists of Interval objects, and not arrays or lists.

// A [[1,4],[7,8],[12,19]]
// B [[0,1],[14,15],[18,20]]
// OUTPUT [[0,4],[7,8],[14,15],[18,19]]

//[[0,4],[7,8],[12,19]]
// [[0,10],[14,15],[18,20]]
//[[0,4],[7,8],[14,15],[18,19]]

func intervalIntersection(A [][]int, B [][]int) [][]int {
	ia := 0
	ib := 0
	merged := [][]int{}
	for {
		if ia >= len(A) || ib >= len(B) {
			break
		}
		var a, b []int
		if ia != len(A) {
			a = A[ia]
		}
		if ib != len(B) {
			b = B[ib]
		}

		if a[0] < b[0] {
			if a[1] >= b[1] {
				merged = append(merged, []int{b[0], b[1]})
				ib++
			} else if a[1] >= b[0] {
				merged = append(merged, []int{b[0], a[1]})
				ia++
			} else {
				ia++
			}

		} else if a[0] == b[0] {
			// a and b start together
			if a[1] > b[1] {
				// b contained in a
				merged = append(merged, []int{a[0], b[1]})

				ib++
			} else if a[1] == b[1] {
				// full overlap
				merged = append(merged, []int{a[0], a[1]})
				ia++
				ib++
			} else /* a[1] < b[1] */ {
				// a contained in b
				merged = append(merged, []int{a[0], b[1]})
				ia++
			}
		} else {
			// a > b
			if b[1] >= a[1] {
				// containment: a in b
				merged = append(merged, []int{a[0], a[1]})
				ia++
				if b[1] == a[1] {
					ib++
				}
			} else if b[1] >= a[0] {
				merged = append(merged, []int{a[0], b[1]})
				ib++
			} else {
				// no overlap
				ib++
			}
		}
	}
	return merged
}

func main() {
	fmt.Println(intervalIntersection(
		[][]int{
			{0, 4},
			{7, 8},
			{12, 19},
		},
		[][]int{
			{0, 10},
			{14, 15},
			{18, 20},
		},
	))
}
