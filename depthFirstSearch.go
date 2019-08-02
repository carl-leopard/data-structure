package main

import "fmt"

type BinaryNode struct {
	val   int
	left  *BinaryNode
	right *BinaryNode
}

type List struct {
	node *BinaryNode
	next *List
}

func main() {
	root := &BinaryNode{
		val: 0,
	}

	root = initTree(root)

	list := dfs(root)
	fmt.Println()
	showList(list)
}

func dfs(root *BinaryNode) *List {
	result := new(List)

	if root == nil || root.left == nil && root.right == nil {
		l := &List{
			node: root,
		}
		result.next = l

		return result
	}

	var stack []*BinaryNode
	var cur *BinaryNode
	visitM := make(map[*BinaryNode]bool)
	head := result

	stack = append(stack, root)
	cnt := 1
	for len(stack) > 0 {
		fmt.Printf("\n %d: ", cnt)
		printStack(stack)

		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if !visitM[cur] {
			visitM[cur] = true

			if cur.right != nil {
				stack = append(stack, cur.right)
			}

			if cur.left != nil {
				stack = append(stack, cur.left)
			}

			l := &List{
				node: cur,
			}
			result.next = l
			result = result.next
		}
		showList(head)
		cnt++

	}

	return head
}

func initTree(root *BinaryNode) *BinaryNode {
	head := root

	root.left = &BinaryNode{
		val: 1,
		left: &BinaryNode{
			val: 3,
			left: &BinaryNode{
				val: 7,
			},
			right: &BinaryNode{
				val: 11,
			},
		},
		right: &BinaryNode{
			val: 4,
			left: &BinaryNode{
				val: 8,
			},
		},
	}
	root.right = &BinaryNode{
		val: 2,
		left: &BinaryNode{
			val: 5,
			left: &BinaryNode{
				val: 9,
			},
			right: &BinaryNode{
				val: 12,
			},
		},
		right: &BinaryNode{
			val: 6,
			right: &BinaryNode{
				val: 10,
			},
		},
	}

	return head
}

func showList(l *List) {
	head := l.next

	fmt.Printf("output: head")
	for head != nil {
		fmt.Printf("->%d", head.node.val)
		head = head.next
	}
}

func printStack(stack []*BinaryNode) {
	fmt.Printf("stack: ")
	for _, n := range stack {
		fmt.Printf("%d->", n.val)
	}
	fmt.Printf("null ")
}
