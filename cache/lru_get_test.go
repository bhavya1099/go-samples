package cache_test

import (
	"testing"

	"github.com/TheAlgorithms/Go/structure/linkedlist"
	"github.com/TheAlgorithms/Go/cache"
)

func TestGet(t *testing.T) {
	// Scenario 1: Retrieve Existing Key Value
	t.Run("Retrieve Existing Key Value", func(t *testing.T) {
		// Arrange
		lru := &cache.LRU{
			Dl:       linkedlist.NewDoublyLinkedList(),
			Size:     0,
			Capacity: 10,
			Storage:  make(map[string]*linkedlist.Node),
		}
		key := "existing_key"
		value := "existing_value"
		node := &linkedlist.Node{Val: cache.Item{Value: value}}

		lru.Storage[key] = node
		lru.Dl.PushFront(node)

		// Act
		retVal := lru.Get(key)

		// Assert
		if retVal != value {
			t.Errorf("Scenario 1: Expected %v, Got %v", value, retVal)
		}
	})

	// Scenario 2: Retrieve Non-Existent Key Value
	t.Run("Retrieve Non-Existent Key Value", func(t *testing.T) {
		// Arrange
		lru := &cache.LRU{
			Dl:       linkedlist.NewDoublyLinkedList(),
			Size:     0,
			Capacity: 10,
			Storage:  make(map[string]*linkedlist.Node),
		}
		key := "non_existent_key"

		// Act
		retVal := lru.Get(key)

		// Assert
		if retVal != nil {
			t.Errorf("Scenario 2: Expected nil, Got %v", retVal)
		}
	})

	// Scenario 3: Move Recently Accessed Key to Back
	t.Run("Move Recently Accessed Key to Back", func(t *testing.T) {
		// Arrange
		lru := &cache.LRU{
			Dl:       linkedlist.NewDoublyLinkedList(),
			Size:     0,
			Capacity: 10,
			Storage:  make(map[string]*linkedlist.Node),
		}
		key := "existing_key"
		value := "existing_value"
		node := &linkedlist.Node{Val: cache.Item{Value: value}}

		lru.Storage[key] = node
		lru.Dl.PushFront(node)

		// Act
		retVal := lru.Get(key)

		// Assert
		if lru.Dl.Back() != node {
			t.Errorf("Scenario 3: Key not moved to the back")
		}
	})

	// Scenario 4: Capacity Limit Handling
	t.Run("Capacity Limit Handling", func(t *testing.T) {
		// Arrange
		lru := &cache.LRU{
			Dl:       linkedlist.NewDoublyLinkedList(),
			Size:     0,
			Capacity: 2, // Low capacity for testing
			Storage:  make(map[string]*linkedlist.Node),
		}
		key1 := "key1"
		value1 := "value1"
		node1 := &linkedlist.Node{Val: cache.Item{Value: value1}}
		key2 := "key2"
		value2 := "value2"
		node2 := &linkedlist.Node{Val: cache.Item{Value: value2}}

		lru.Storage[key1] = node1
		lru.Dl.PushFront(node1)
		lru.Storage[key2] = node2
		lru.Dl.PushFront(node2)

		// Act
		_ = lru.Get(key1)

		// Assert
		if lru.Dl.Size() > lru.Capacity {
			t.Errorf("Scenario 4: Cache exceeded capacity")
		}
	})

	// Scenario 5: Concurrent Access
	t.Run("Concurrent Access", func(t *testing.T) {
		// TODO: Implement concurrent access test scenario
		t.Skip("Scenario 5: Concurrent Access test not implemented")
	})
}
