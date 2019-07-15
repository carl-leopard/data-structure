package main

import "fmt"


 // Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

func main() {
	head := &ListNode{0, nil}
	h := head
	for i := 1; i< 6; i++ {
		h.Next = &ListNode{i, nil}
		h = h.Next
	}
	
	show(head)
	
	head = reverseKGroup(head, 1)
	show(head)
	
	head = reverseKGroup(head, 2)
	show(head)
	
	head = reverseKGroup(head, 3)
	show(head)
	
	head = reverseKGroup(head, 4)
	show(head)
	
	head = reverseKGroup(head, 5)
	show(head)
	
	head = reverseKGroup(head, 6)
	show(head)
	
	head = reverseKGroup(head, 7)
	show(head)
	
}
 
 
 /*
 翻转之后， K group 中的头节点会变为尾节点， 尾节点会变为头节点
 	1. 根据 K， 查找需要翻转的 listNode 的 head， tail
 	2. 记录上一个翻转的 K group 的尾节点， 用于将各个 K group 串联起来
 	3. result 用于表示翻转之后返回的头节点： 第一次翻转得到的节点就是新的链表的头节点
	4. 当剩余的节点数小于 K 时， 不需要翻转
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if k <= 1 {
		return head
	}
	
	
    iterator := head
    newHead := head
    
    var (
            h *ListNode
            tail *ListNode
            lastTail *ListNode
            
            result *ListNode
        )
    
    for  {
        tail = nil
        
        for i := 1; i <= k && iterator != nil; i++ {
            if i == 1 {
                h = iterator
            }
            
            if i == k {
                tail = iterator
            }
            
            iterator = iterator.Next
        }
        
        if tail != nil {
        	if lastTail != nil {
        		lastTail.Next = tail
        	}
        	
        	lastTail = h
        	
        	result = reverseByRange(h, tail)
        	
        	if newHead == head {
        		newHead = result
        	}
        	
        	fmt.Printf("newHead:%d, h:%d, tail:%d, result:%d, lastTail:%d\n", newHead.Val, h.Val, tail.Val, result.Val, lastTail.Val)
        	show(newHead)
        } else {
        	break
        }
    }
    
    return newHead
}


//reverse listNode by range (head, tail)
func reverseByRange(head, tail *ListNode) *ListNode{   
    var (
        pre *ListNode
        next *ListNode
    )
    
    if tail != nil {
       pre = tail.Next
    }
    for head != nil && head != tail{
        next = head.Next
        
        head.Next = pre
        pre = head
        head = next
    }
    
    if head != nil {
		head.Next = pre
    	pre = head
    }
    
    return pre
}


//show listNode
func show(head *ListNode) {
	tmp := head
	for tmp != nil {
		fmt.Printf("%d->", tmp.Val)
		tmp = tmp.Next
	}
	fmt.Println("NULL")
}