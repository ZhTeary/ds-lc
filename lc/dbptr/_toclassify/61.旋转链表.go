package doublePointer

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	//判断空链表
	if head == nil {
		return head
	}
	//先遍历一遍找到链表长度
	point := head
	count := 0
	for point != nil {
		count++
		point = point.Next
	}

	//再算出来该第几个几点当头
	new := count - (k % count) + 1

	//将首尾相连
	tail := head
	for i := 0; i < count-1; i++ {
		tail = tail.Next
	}
	tail.Next = head

	//定位新头并将新头前面断开
	pre := head
	for i := 0; i < new-2; i++ {
		pre = pre.Next
	}

	newhead := pre.Next
	pre.Next = nil

	return newhead
}
