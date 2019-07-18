package main

import "fmt"

//https://leetcode.com/problems/reverse-nodes-in-k-group/

 // Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

func main() {
	head := &ListNode{0, nil}
	h := head
	for i := 1; i< 2; i++ {
		h.Next = &ListNode{i, nil}
		h = h.Next
	}
	
	
	reverseKGroup(head, 2)
	
}
 
 
 /*
 翻转之后， K group 中的头节点会变为尾节点， 尾节点会变为头节点
 	1. 根据 K， 查找需要翻转的 listNode 的 head， tail
 	2. 记录上一个翻转的 K group 的尾节点， 用于将各个 K group 串联起来
 	3. result 用于表示翻转之后返回的头节点： 第一次翻转得到的节点就是新的链表的头节点
	4. 当剩余的节点数小于 K 时， 不需要翻转
 */
func reverseKGroup(head *ListNode, k int)  *ListNode{
	if head == nil || k <= 1 {
		return head
	}
	
	nodes := make([]*ListNode, 0, 10)
	var next *ListNode
	for head != nil {
		next = head.Next
		head.Next = nil
		nodes = append(nodes, head)
		head = next
	}
	
	reverseByArray(nodes, k)
	
	show(nodes)
	
	
	p, iterator := nodes[0], nodes[0]
    for i:=1; i<len(nodes); i++ {
		iterator.Next = nodes[i]
		iterator = iterator.Next
	}
	
	showList(p)
	
	return p
}

func reverseByArray(nodes []*ListNode, k int)  {
	if k <= 1 {
		return
	}
	
	if len(nodes)  < k {
		return
	}
	
	var (
		start int
		end int
	)
	
	end = -1
	for {
		start = end + 1
		end = k - 1 + start 
		
		if end >= len(nodes) {
			break
		}
		
		reverseByRange(nodes, start, end)
		
		show(nodes)
	}
}


//reverse listNode by range (head, tail)
func reverseByRange(nodes []*ListNode, start, end int){   
    for start < end {
    	nodes[start], nodes[end] = nodes[end], nodes[start]
    	
    	start++
    	end--
    	
    }
}


//show listNode
func show(nodes []*ListNode) {
	for _, node := range nodes{
		fmt.Printf("%d->", node.Val)
		node = node.Next
	}
	fmt.Println("NULL")
}

//show listNode
func showList(head *ListNode) {
	tmp := head
	for tmp != nil {
		fmt.Printf("%d->", tmp.Val)
		tmp = tmp.Next
	}
	fmt.Println("NULL")
}