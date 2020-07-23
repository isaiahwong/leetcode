package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	zz := [][]int{}
	if root == nil {
		return zz
	}
	row := []int{}
	queue := []*TreeNode{}
	queue = append(queue, root)
	// sets the first level
	queue = append(queue, nil)
	// true indicates the direction going to the right
	cross := true

	for len(queue) > 0 {
		currentNode := queue[0]
		// Remove node
		queue = queue[1:]

		// keep tracks of level with nil
		if currentNode == nil {
			// Add same level
			zz = append(zz, row)
			// Reset row
			row = []int{}
			queue = append(queue, nil)
			if queue[0] == nil {
				break
			} else {
				// Change direction
				cross = !cross
				continue
			}
		}
		if currentNode.Left != nil {
			queue = append(queue, currentNode.Left)
		}
		if currentNode.Right != nil {
			queue = append(queue, currentNode.Right)
		}
		if cross {
			// We append if direction is going to the right
			row = append(row, currentNode.Val)
		} else {
			// we prepend if direction is going to the left
			row = append([]int{currentNode.Val}, row...)
		}
	}
	return zz
}

func main() {
	// Test cases
	nodes := []*TreeNode{
		{
			Val: 3,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val:   20,
				Left:  &TreeNode{Val: 15},
				Right: &TreeNode{Val: 7},
			},
		},
		{
			Val: 3,
			Left: &TreeNode{
				Val:   9,
				Left:  &TreeNode{Val: 222},
				Right: &TreeNode{Val: 2},
			},
			Right: &TreeNode{
				Val:   20,
				Left:  &TreeNode{Val: 15},
				Right: &TreeNode{Val: 7},
			},
		},
		nil,
		{
			Val: 3,
			Left: &TreeNode{
				Val: 9,
				Left: &TreeNode{
					Val:  12,
					Left: &TreeNode{Val: 9},
					Right: &TreeNode{Val: 2,
						Left:  &TreeNode{Val: 90},
						Right: &TreeNode{Val: 23},
					},
				},
				Right: &TreeNode{Val: 8},
			},
			Right: &TreeNode{
				Val:   20,
				Left:  &TreeNode{Val: 15},
				Right: &TreeNode{Val: 7},
			},
		},
	}
	for _, n := range nodes {
		fmt.Println(zigzagLevelOrder(n))
	}
}
