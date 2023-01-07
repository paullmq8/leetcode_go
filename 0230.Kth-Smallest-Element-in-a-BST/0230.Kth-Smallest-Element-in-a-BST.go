package main

import "fmt"

/*
type TreeNode = structures.TreeNode

func kthSmallest(root *TreeNode, k int) int {
    strings.Builder
}
*/

func main() {
	arr := []int{}
	recursion(&arr, 1)
	recursion(&arr, 2)
	fmt.Println(arr)
}

func recursion(arr *[]int, i int) {
	*arr = append(*arr, i)
}
