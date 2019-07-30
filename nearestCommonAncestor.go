package main

func nearesrCommonAncestor(root *BinaryTree, p *BinaryTree, q *BinaryTree) *BinaryTree {
	return findPorQ(root, p, q)
}

func findPorQ(root *BinaryTree, p *BinaryTree, q *BinaryTree) *BinaryTree {
	//查找到叶子节点
	if root == nil {
		return root
	}

	//跟节点是p/q, 则公共父节点即为根节点
	if root == p || root == q {
		return root
	}

	left := findPorQ(root.left, p, q)
	right := findPorQ(root.right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left == nil && right == nil {
		return root
	}

	if right != nil {
		return left
	}

	return left
}
