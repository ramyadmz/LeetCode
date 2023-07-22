package lrucache

type Node struct {
	key, value int
	next, prev *Node
}
type LRUCache struct {
	keyToNode  map[int]*Node
	head, tail *Node
	capacity   int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		keyToNode: make(map[int]*Node),
		capacity:  capacity,
	}
}

// AddNode creates a new node with the given key and value, and adds it to the cache.
func (this *LRUCache) addNode(key, value int) *Node {
	node := &Node{
		key:   key,
		value: value,
		next:  this.tail,
	}
	if this.head == nil {
		this.head = node
	}
	if this.tail != nil {
		this.tail.prev = node
	}
	this.tail = node
	return node
}

// RemoveNode removes the given node from the cache.
func (this *LRUCache) removeNode(node *Node) {
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		this.tail = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	} else {
		this.head = node.prev
	}
}

func (this *LRUCache) Get(key int) int {

}

func (this *LRUCache) Put(key int, value int) {

}
