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

// Get retrieves the value of the given key if it exists in the cache, else returns -1.
func (this *LRUCache) Get(key int) int {
	node, exist := this.keyToNode[key]
	if !exist {
		return -1
	}
	this.removeNode(node)
	newNode := this.addNode(key, node.value)
	this.keyToNode[key] = newNode
	return node.value
}

// Put adds a new node with the given key and value to the cache.
func (this *LRUCache) Put(key int, value int) {
	if node, exists := this.keyToNode[key]; exists {
		this.removeNode(node)
	} else if this.capacity == len(this.keyToNode) {
		delete(this.keyToNode, this.head.key)
		this.removeNode(this.head)
	}
	newNode := this.addNode(key, value)
	this.keyToNode[key] = newNode
}
