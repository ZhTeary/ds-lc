/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

package linkedList

type LRUCache struct {
	size       int
	capacity   int
	data       map[int]*DLinkedNode
	head, tail *DLinkedNode
}
type DLinkedNode struct {
	key   int
	value int
	pre   *DLinkedNode
	next  *DLinkedNode
}

func NewDLinkedNode(key int, value int) *DLinkedNode {
	return &DLinkedNode{key: key, value: value}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		size:     0,
		capacity: capacity,
		data:     make(map[int]*DLinkedNode, capacity),
		head:     NewDLinkedNode(0, 0),
		tail:     NewDLinkedNode(0, 0),
	}
	l.head.next = l.tail
	l.tail.pre = l.head
	return l

}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.data[key]; !ok {
		return -1
	}
	node := this.data[key]
	this.MoveToHead(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	/*
		if _, ok := this.data[key]; !ok {
			node := NewDLinkedNode(key, value)
			this.data[key] = node
			this.addToHead(node)
			this.size++
			if this.size > this.capacity {
				removed := this.tail.pre
				removed.pre.next = removed.next
				removed.next.pre = removed.pre
				delete(this.data, removed.key)
				this.size--
			}
		} else {
			node := this.data[key]
			node.value = value
			this.moveToHead(node)
		}
	*/
	//已经存在
	if _, ok := this.data[key]; !ok {
		node := NewDLinkedNode(key, value)
		this.data[key] = node
		this.addToHead(node)
		this.size++
		//超了
		if this.size > this.capacity {
			removed := this.tail.pre
			removed.pre.next = removed.next
			removed.next.pre = removed.pre
			//this.removeNode(removed)
			//this.data[removed.key] = nil
			delete(this.data, removed.key)

			this.size--
		}
	} else {
		node := this.data[key]
		node.value = value
		this.MoveToHead(node)
	}
	//不存在
	//插入

}

func (this *LRUCache) MoveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}
func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
}
func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.pre = this.head
	node.next = this.head.next
	this.head.next.pre = node
	this.head.next = node
}
