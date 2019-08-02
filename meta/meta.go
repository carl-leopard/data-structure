package meta

type BinaryNode struct {
	Val        int
	LeftChild  *BinaryNode
	RightChild *BinaryNode
}

type MultiWayNode struct {
	Val      int
	Children []*MultiWayNode
}

type AVLNode struct {
	Val        int
	LeftChild  *AVLNode
	RightChild *AVLNode

	Depth  int //left-right=1|0|-1
	Parent *AVLNode
}
