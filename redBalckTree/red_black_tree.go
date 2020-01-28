package main

import "fmt"

const (
	RedBlackNode_Color_Red = iota
	RedBlackNode_Color_Black
)

type RedBlackNode struct {
	Lchild *RedBlackNode
	Rchild *RedBlackNode
	Parent *RedBlackNode
	Val    int
	Color  int
}

func main() {
	array := []int{100, 50, 200, 2, 3, 6}
	//array := []int{100, 50, 200, 2, 3}
	if len(array) == 0 {
		return
	}
	root := &RedBlackNode{
		Val:   array[0],
		Color: RedBlackNode_Color_Black,
	}
	MidOrder(root)
	fmt.Println()

	for i := 1; i < len(array); i++ {
		node := newRedBlackNode(array[i])
		Insert(root, node)

		MidOrder(root)
		fmt.Println()
	}

	MidOrder(root)

	fmt.Println("\npreOrder:")
	PreOrder(root)

	fmt.Println("\npostOrder:")
}

func Insert(root, node *RedBlackNode) {
	if root == nil {
		root = node
	} else if node.Val > root.Val {
		if root.Rchild == nil {
			root.Rchild = node
			node.Parent = root
		} else {
			Insert(root.Rchild, node)
		}
	} else {
		if root.Lchild == nil {
			root.Lchild = node
			node.Parent = root
		} else {
			Insert(root.Lchild, node)
		}
	}

	root.adjust(node)
}

func (root *RedBlackNode) adjust(node *RedBlackNode) {
	var (
		parent, uncle *RedBlackNode
	)

	//插入的节点是根节点
	if node.Parent == nil {
		root.insertRoot(node)
		return
	}

	//插入节点的父节点是黑色的
	parent = node.Parent
	if parent.Color == RedBlackNode_Color_Black {
		root.insertParentIsBlack(node)
		return
	}

	//以下情况，父节点是红色的， 祖父节点是黑色的（性质4）

	//祖父节点是黑色的， 父节点是红色的，叔父节点也是红色的
	uncle = getUncle(node)
	if parent.Color == RedBlackNode_Color_Red && uncle != nil && uncle.Color == RedBlackNode_Color_Red {
		root.insertParentAndUncleAreRed(node)
	}

	//祖父节点是黑色的， 父节点是红色的，叔父节点是黑色的或者空节点，当前节点是父节点的右孩子
	if parent.Color == RedBlackNode_Color_Red &&
		(uncle == nil || uncle.Color == RedBlackNode_Color_Black) &&
		node == parent.Rchild {
		root.insertPrUbnNright(node)
	}

	//祖父节点是黑色的， 父节点是红色的，叔父节点是黑色的或者空节点，当前节点是父节点的左孩子
	if parent.Color == RedBlackNode_Color_Red &&
		(uncle == nil || uncle.Color == RedBlackNode_Color_Black) &&
		node == parent.Lchild {
		root.insertPrUbnNleft(node)
	}
}

func MidOrder(root *RedBlackNode) {
	if root == nil {
		return
	}

	MidOrder(root.Lchild)

	color := "red"
	if root.Color == RedBlackNode_Color_Black {
		color = "black"
	}
	fmt.Printf("(%d:%s) ", root.Val, color)

	MidOrder(root.Rchild)
}

func PreOrder(root *RedBlackNode) {
	if root == nil {
		return
	}

	color := "red"
	if root.Color == RedBlackNode_Color_Black {
		color = "black"
	}
	fmt.Printf("(%d:%s) ", root.Val, color)

	PreOrder(root.Lchild)
	PreOrder(root.Rchild)
}

//1. case 1 新节点 位于根节点，没有父节点。性质2. 性质5
func (root *RedBlackNode) insertRoot(node *RedBlackNode) {
	node.Color = RedBlackNode_Color_Black
}

//2. case 2 新增节点的父节点是黑色的，直接插入即可
func (root *RedBlackNode) insertParentIsBlack(node *RedBlackNode) {
	return
}

/*
以下场景认为:
	1. 新增节点的父节点都是红色的 -> 一定有祖父节点
	2. 如果祖父节点是根节点 -> 祖父节点应该是黑色的，所以新节点一定有个叔父节点
*/

//3.
/* case 3
1. 新增节点的父节点和叔节点都是红色 -> 祖父节点一定是黑色的
*/
/*solution
1. 将父节点叔父节点置为黑色
2. 将祖父节点置为红色
3. 祖父节点可能是根节点违反了性质2，或者祖父节点的父节点是红色的违反性质4： 对祖父节点进行新插入的操作
*/

func (root *RedBlackNode) insertParentAndUncleAreRed(node *RedBlackNode) {

	//性质5
	parent := node.Parent
	parent.Color = RedBlackNode_Color_Black
	uncle := getUncle(node)
	uncle.Color = RedBlackNode_Color_Black
	grand := getGrandParent(node)
	grand.Color = RedBlackNode_Color_Red

	//红色的祖父节点有可能是根节点
	node = grand
	root.adjust(node)
}

//假定父节点是祖父节点的左孩子
/* case 4
1. 父节点是红色的 -> 祖父节点是黑色的
2. 叔父节点是黑色的或者缺失
3. 新节点是父节点的右孩子
*/

/* solution
1. 对父节点进行一次左旋
2. DO case PrUbnNl
*/
func (root *RedBlackNode) insertPrUbnNright(node *RedBlackNode) {
	grand := getGrandParent(node)
	parent := node.Parent

	if parent == grand.Lchild && node == parent.Rchild {
		grand.Lchild = rotateLeft(parent) // turn to case insertPrUbnNl
	} else if parent == grand.Rchild && node == parent.Rchild {
		grand.Rchild = rotateRight(parent)
	}

	node = parent
	root.adjust(node)
}

/* case 5
1. 父节点是红色的
2. 叔父节点是黑色的或者缺失
3. 新节点是父节点的左孩子
*/

/* solution
1. 将父节点置位黑色，祖父节点置位红色
2. 对祖父节点右旋
3. 父节点是红色的，祖父节点一定是黑色的
*/
func (root *RedBlackNode) insertPrUbnNleft(node *RedBlackNode) {
	parent := node.Parent
	parent.Color = RedBlackNode_Color_Black
	grand := getGrandParent(node)
	grand.Color = RedBlackNode_Color_Red

	parentOfGrand := grand.Parent
	var n *RedBlackNode
	if node == parent.Lchild && parent == grand.Lchild {
		n = rotateRight(grand)
	} else {
		//node = parent.Lchild && parent = gand.Rchild
		n = rotateLeft(grand)
	}

	if parentOfGrand == nil {
		root = n
	} else if grand == parentOfGrand.Lchild {
		parentOfGrand.Lchild = n
	} else {
		parentOfGrand.Rchild = n
	}
}

func newRedBlackNode(val int) *RedBlackNode {
	return &RedBlackNode{
		Val:   val,
		Color: RedBlackNode_Color_Red,
	}
}

func getGrandParent(node *RedBlackNode) *RedBlackNode {
	return node.Parent.Parent
}

func getUncle(node *RedBlackNode) *RedBlackNode {
	grandParent := getGrandParent(node)
	parent := node.Parent
	if parent == grandParent.Lchild {
		return grandParent.Rchild
	}

	return grandParent.Lchild
}

func rotateLeft(node *RedBlackNode) *RedBlackNode {
	rchild := node.Rchild

	node.Rchild = rchild.Lchild
	rchild.Lchild = node

	rchild.Parent = node.Parent
	node.Parent = rchild
	return rchild
}

func rotateRight(node *RedBlackNode) *RedBlackNode {
	lchild := node.Lchild

	node.Lchild = lchild.Rchild
	lchild.Rchild = node

	lchild.Parent = node.Parent
	node.Parent = lchild
	return lchild
}
