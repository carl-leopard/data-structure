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
