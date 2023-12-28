package linkedList

type ListNode struct {
    Val int
    Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	hashmap := make(map[ListNode]bool)
	
	temp := head
	for temp != nil{
		if _,ok := hashmap[*temp];ok{
			return temp
		}
		hashmap[*temp] = true
		temp = temp.Next
		}

	return nil
}