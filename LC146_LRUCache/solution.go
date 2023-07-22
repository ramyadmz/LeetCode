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

func (this *LRUCache) Get(key int) int {

}

func (this *LRUCache) Put(key int, value int) {

}
