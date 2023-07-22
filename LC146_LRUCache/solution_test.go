package lrucache

import (
	"testing"
)

func TestLRUCache(t *testing.T) {
	capacity := 3
	cache := Constructor(capacity)

	// Test adding elements to the cache
	cache.Put(1, 10)
	cache.Put(2, 20)
	cache.Put(3, 30)
	if cache.Get(1) != 10 {
		t.Errorf("Expected cache.Get(1) to return 10")
	}

	// Test adding another element that exceeds the cache capacity
	cache.Put(4, 40)
	if cache.Get(2) != -1 {
		t.Errorf("Expected cache.Get(2) to return -1 after adding a new element and exceeding capacity")
	}

	// Test updating an existing element
	cache.Put(3, 35)
	if cache.Get(3) != 35 {
		t.Errorf("Expected cache.Get(3) to return 35 after updating the value")
	}

	// Test removing an element from the cache
	cache.Put(5, 50)
	if cache.Get(1) != -1 {
		t.Errorf("Expected cache.Get(1) to return -1 after removing the element")
	}
}
