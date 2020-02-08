package main

import "fmt"

const (
	M           int = 3
	MaxKeyNum       = M - 1
	MinKeyNum       = (M+1)/2 - 1
	MidKeyIndex     = M / 2
	MaxChildNum     = M
)

type BPlusNode struct {
	keys   []int
	keyNum int
	child  []*BPlusNode
	parent *BPlusNode
}

func main() {
	array := []int{100, 200, 10, 89, 65, 32, 17, 27, 99, 123, 500, 70}
	var root *BPlusNode
	for _, v := range array {
		root = Insert(root, v)
		BrandReverser(root)
		fmt.Println()
		InOrder(root)
		fmt.Println()
		fmt.Println()
	}
}

func newBPlusNode(val []int) *BPlusNode {
	node := &BPlusNode{
		keys:   make([]int, 0, M),
		keyNum: len(val),
		child:  make([]*BPlusNode, 0, MaxChildNum),
		parent: nil,
	}

	node.keys = append(node.keys, val...)

	return node
}

func Insert(root *BPlusNode, val int) *BPlusNode {
	node := searchPos(root, val)

	if node == nil {
		newNode := newBPlusNode([]int{val})
		return newNode
	}

	node.keys, _ = sliceInsert(node.keys, val)
	node.keyNum++

	//judge keyNum
	if node.keyNum <= MaxKeyNum {
		return root
	}

	newNode := split(node)
	if newNode == nil {
		return root
	}

	return newNode
}

func BrandReverser(root *BPlusNode) {
	if root == nil {
		return
	}

	for _, v := range root.keys {
		fmt.Printf("%d,", v)
	}
	fmt.Printf("  ")

	for _, child := range root.child {
		BrandReverser(child)
	}
}

func InOrder(root *BPlusNode) {
	if root == nil {
		return
	}

	if len(root.child) == 0 {
		for _, v := range root.keys {
			fmt.Printf("%d ", v)
		}
		return
	}

	for i := 0; i < root.keyNum; i++ {
		if len(root.child) > i {
			InOrder(root.child[i])
		}
		fmt.Printf("%d ", root.keys[i])
	}

	if len(root.child) == root.keyNum+1 {
		InOrder(root.child[root.keyNum])
	}
}

func searchPos(root *BPlusNode, val int) *BPlusNode {
	//tree is empty
	if root == nil {
		return root
	}

	index := 0
	for ; index < root.keyNum && val > root.keys[index]; index++ {
	}

	//no child
	if len(root.child) == 0 || root.child[index] == nil {
		return root
	}

	return searchPos(root.child[index], val)
}

func split(node *BPlusNode) *BPlusNode {
	left := newBPlusNode(getByRange(node.keys, 0, MidKeyIndex-1))
	right := newBPlusNode(getByRange(node.keys, MidKeyIndex+1, node.keyNum-1))
	parent := node.parent
	if parent == nil {
		parent = newBPlusNode([]int{node.keys[MidKeyIndex]})
		parent.child = append(parent.child, left, right)
		setParentForChild(parent)
		reDispatchChild(node, left, right, node.keys[MidKeyIndex])
		return parent
	}

	var index int
	parent.keys, index = sliceInsert(parent.keys, node.keys[MidKeyIndex])
	parent.keyNum++

	r := append([]*BPlusNode{right}, parent.child[index+1:]...)
	parent.child = append(parent.child[:index], left)
	parent.child = append(parent.child, r...)
	setParentForChild(parent)
	reDispatchChild(node, left, right, node.keys[MidKeyIndex])
	if parent.keyNum <= MaxKeyNum {
		return nil
	}

	return split(parent)
}

func setParentForChild(root *BPlusNode) {
	if root == nil {
		return
	}

	for i := 0; i < len(root.child); i++ {
		root.child[i].parent = root
	}
}

func reDispatchChild(node, left, right *BPlusNode, key int) {
	if len(node.child) == 0 {
		return
	}

	for _, child := range node.child {
		if child.keys[len(child.keys)-1] < key {
			left.child = append(left.child, child)
		} else {
			right.child = append(right.child, child)
		}
	}
	setParentForChild(left)
	setParentForChild(right)
}

func sliceInsert(keys []int, val int) ([]int, int) {
	index := 0
	for ; index < len(keys) && val > keys[index]; index++ {
	}

	return sliceInsertByIndex(keys, val, index), index

}

func sliceInsertByIndex(keys []int, val, index int) []int {
	if index >= len(keys) {
		keys = append(keys, val)
		return keys
	}

	//keys = append([]int{val}, keys...)
	if index == 0 {
		arr := make([]int, 0, len(keys)+1)
		arr = append(arr, val)
		arr = append(arr, keys...)
		return arr
	}

	//keys = append(keys[:index], append([]int{val}, keys[index:]))
	arr := make([]int, len(keys)-index)
	copy(arr, keys[index:])
	keys = append(keys[:index], val)
	keys = append(keys, arr...)
	return keys
}

func getByRange(keys []int, start, end int) []int {
	if len(keys) == 0 ||
		start > end ||
		start < 0 ||
		start >= len(keys) ||
		end >= len(keys) {
		return make([]int, 0, 0)
	}

	dst := make([]int, end-start, end-start+1)
	copy(dst, keys[start:end])
	dst = append(dst, keys[end])
	return dst
}
