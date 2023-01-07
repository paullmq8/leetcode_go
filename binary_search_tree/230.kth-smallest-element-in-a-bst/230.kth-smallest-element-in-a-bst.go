package binary_search_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * @lc app=leetcode id=230 lang=golang
 *
 * [230] Kth Smallest Element in a BST
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
	res := -1
	inOrderDfs(root, &k, &res)
	return res
}

func inOrderDfs(node *TreeNode, k *int, res *int) {
	if node == nil {
		return
	}
	inOrderDfs(node.Left, k, res)
	if *k == 1 {
		*res = node.Val
	}
	*k--
	inOrderDfs(node.Right, k, res)
}

// @lc code=end
