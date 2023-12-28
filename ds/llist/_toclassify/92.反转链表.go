package linkedList
// /**
//  * Definition for singly-linked list.
//  * type ListNode struct {
//  *     Val int
//  *     Next *ListNode
//  * }
//  */


// type ListNode struct {
// 	     Val int
// 	     Next *ListNode
// }

 
// func reverseBetween(head *ListNode, left int, right int) *ListNode {
// 	//反转位置相等与空链表
// 	if left == right || head == nil{
// 		return head
// 	}

// 	//链表长度为2 
// 	len := 0
// 	temptr:= head
// 	for temptr != nil{
// 		len++ 
// 		temptr = temptr.Next
// 	}
// 	if len == 2 {
// 		temp := head.Next
// 		temp.Next = head
// 		head.Next = nil
// 		return temp
// 	}
	
// 	//定位到left的前一个节点
// 	i := 0
// 	var PreL = &ListNode{Next:head}
// 	for i < left-1 {
// 		PreL = PreL.Next
// 		i++
// 	}
	
// 	//标记初始指针
// 	L := PreL.Next
// 	M := L.Next
// 	R := M.Next
// 	//为最后一次寻找头尾做打算
// 	Last := L
// 	//循环反转链表
// 	l := left
// 	for l < right {
// 		M.Next = L
// 		L = M 
// 		M = R 
// 		if R == nil{
// 			break
// 		}else{
// 		R = R.Next
// 		}
// 		l++
// 	}
// 	PreL.Next = L
// 	Last.Next = M

// 	if left == 1 {
// 		return PreL.Next
// 	}else{
// 			return head
// 	}

// } 
