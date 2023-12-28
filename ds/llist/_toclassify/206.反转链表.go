package linkedList





// type ListNode struct {
//     Val int
//     Next *ListNode
// }

// //2023-11-13 第二次做
// func reverseList(head *ListNode) *ListNode {
// 	//链表为空或单个节点
// 	if head == nil || head.Next==nil{
// 		return head
// 	}

// 	pre := head
// 	mid := pre.Next
// 	nxt := mid.Next
//     pre.Next = nil

// 	for nxt != nil{
// 		mid.Next = pre
// 		pre = mid
// 		mid = nxt
// 		nxt = nxt.Next
// 	}
// 	mid.Next = pre

// 	return mid
// }



// 简单
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
//func reverseList(head *ListNode) *ListNode {
//	//双指针循环一遍操作
//
//	//特殊情况
//	if head == nil { //没有节点
//		return nil
//	}
//	if head.Next == nil { // 一个节点
//		return head
//	}
//	if head.Next.Next == nil { //两个节点
//		pre := head
//		next := pre.Next
//		next.Next = pre
//		pre.Next = nil
//		return next
//	}
//
//	pre := head
//	mid := head.Next
//	nxt := mid.Next
//	head.Next = nil
//
//	for nxt != nil {
//		mid.Next = pre
//		pre = mid
//		mid = nxt
//		nxt = nxt.Next
//	}
//	mid.Next = pre
//
//	return mid
//
//}
