package linkedList

//
////中等
//
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
//func deleteMiddle(head *ListNode) *ListNode {
//	//get-linkedlenth
//	listlenth := 0
//	ptr := head
//	for ptr != nil {
//		listlenth++
//		ptr = ptr.Next
//	}
//	if listlenth == 1 {
//		return nil
//	}
//
//	left := head
//	for i := 0; i < (listlenth/2)-1; i++ {
//		left = left.Next
//	}
//
//	left.Next = left.Next.Next
//
//	return head
//}
