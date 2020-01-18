package meta

import (
	"github.com/carl-leopard/data-structure/utils"
)

const (
	BALANCE_FACTOR_LEFT_HIGH  = 2
	BALANCE_FACTOR_RIGHT_HIGH = -2
)

type AVLNode struct {
	Val   int
	Left  *AVLNode
	Right *AVLNode

	Height int
}

func NewAVLNode(val int) *AVLNode {
	return &AVLNode{
		Val:   val,
		Left:  nil,
		Right: nil,

		Height: 1,
	}
}

func (avl *AVLNode) Adjust() {
	balanceFactor := avl.Left.GetHeight() - avl.Right.GetHeight()

	switch balanceFactor {
	case BALANCE_FACTOR_LEFT_HIGH:
		if avl.Left.Left.GetHeight() - avl.Left.Right.GetHeight() {
			avl.RightRotate()
			return
		}

		avl.LRRotate()

	case BALANCE_FACTOR_RIGHT_HIGH:
		if avl.Right.Right.GetHeight() - avl.Right.Left().GetHeight() {
			avl.LeftRotate()
			return
		}

		avl.RLRotate()

	}
}

func (avl *AVLNode) InsertAVL(val int) *AVLNode {
	if avl == nil {
		avl = NewAVLNode(val)
		return
	}

	switch {
	case val < avl.Val:
		avl.Left = avl.Left.InsertAVL(val)
		avl.Adjust()
	case val > avl.Val:
		avl.Right = avl.Right.InsertAVL(val)
		avl.Adjust()
	default:
		fmt.Println("data already exists")
	}

	avl.Height = utils.Max(avl.Left.GetHeight(), avl.Right.GetHeight()) + 1

	return
}

func (avl *AVLNode) DeleteAVL(val int) *AVLNode {
	find := avl.Find(val)
	if find == nil {
		fmt.Println("can not find node")
		return nil
	}

	if find.Left != nil && find.Right != nil {
		find.Val = find.Right.FindMin().Val
		find.Right.DeleteAVL(find.Val)
	} else if find.Left != nil {
		find = find.Left
	} else {
		find = find.Right
	}

	if find != nil {
		find.Height = utils.Max(find.Left.GetHeight(), find.Right.GetHeight()) + 1
		find.Adjust()
	}

	return find
}

func (avl *AVLNode) Find(val int) *AVLNode {
	if avl == nil {
		return nil
	}

	if val == avl.Val {
		return avl
	}

	if val > avl.Val {
		return avl.Right.Find(val)
	}

	return avl.Left.Find(val)
}

func (avl *AVLNode) FindMax() *AVLNode {
	if avl == nil {
		return nil
	}

	if avl.Right == nil {
		return avl
	}

	return avl.Right.FindMax()
}

func (avl *AVLNode) FindMin() *AVLNode {
	if avl == nil {
		return nil
	}

	if avl.Left == nil {
		return avl
	}

	return avl.Left.FindMin()
}

//RightRotate rotates the whole tree towards right
func (avl *AVLNode) RightRotate() *AVLNode {
	head := avl.Left

	avl.Left = head.Right
	head.Right = avl

	avl.Height = utils.Max(avl.Left.GetHeight(), avl.Right.GetHeight()) + 1
	head.Height = utils.Max(head.Left.GetHeight(), head.Right.GetHeight()) + 1

	return head
}

//LeftRotate rotates the whole tree towards left
func (avl *AVLNode) LeftRotate() *AVLNode {
	head := avl.Right

	avl.Right = head.Left
	hread.Left = avl

	avl.Height = utils.Max(avl.Left.GetHeight(), avl.Right.GetHeight()) + 1
	head.Height = utils.Max(head.Left.GetHeight(), head.Right.GetHeight()) + 1

	return head
}

//LRRotate rotates left tree towards left, then rotates the whole tree towards right
func (avl *AVLNode) LRRotate() *AVLNode {
	avl.Left = avl.Left.LeftRotate()

	return avl.RightRotate()
}

//RLRotate rotates right tree towards right, then rotates the whole tree towards left
func (avl *AVLNode) RLRotate() *AVLNode {
	avl.Right = avl.Right.RightRotate()

	return avl.LeftRotate()
}

func (avl *AVLNode) GetHeight() int {
	if avl == nil {
		return 0
	}

	return avl.Height
}
