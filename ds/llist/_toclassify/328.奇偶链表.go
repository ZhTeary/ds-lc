package linkedList

//中等

//O(n)的时间复杂度，O(1)的空间复杂度 ==》 原地解决

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
//func OddEvenList(head *ListNode) *ListNode {
//	//特殊情况
//	if head == nil {
//		return nil
//	}
//	if head.Next == nil {
//		return head
//	}
//
//	//找到结尾
//	tail := head
//	listlenth := 1 // 偶数节点ll/2 ， 奇数节点ll-偶
//	for tail.Next != nil {
//		tail = tail.Next
//		listlenth++
//	}
//	if listlenth == 2 {
//		return head
//	}
//
//	Odd := head
//	for i := 0; i < listlenth/2; i++ {
//		Even := Odd.Next
//		Odd.Next = Even.Next
//		Odd = Even.Next
//		tail.Next = Even
//		Even.Next = nil
//		tail = Even
//	}
//
//	return head
//}
