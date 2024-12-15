package cache

import (
	"testing"
	"github.com/TheAlgorithms/Go/structure/linkedlist"
)

func TestLRUPut(t *testing.T) {
	// Scenario 1: Put existing key with new value
	t.Run("put_existing_key_with_new_value", func(t *testing.T) {
		// Arrange
		lru := &LRU{
			dl:       linkedlist.NewDoubly(),
			size:     1,
			capacity: 3,
			storage:  make(map[string]*linkedlist.Node),
		}
		lru.Put("abc", 123)

		// Act
		lru.Put("abc", 456)

		// Assert
		node := lru.storage["abc"]
		if node.Val.(item).value != 456 {
			t.Errorf("Scenario 1: Value for key 'abc' not updated correctly")
		}
	})

	// Scenario 2: Put new key in full cache
	t.Run("put_new_key_in_full_cache", func(t *testing.T) {
		// Arrange
		lru := &LRU{
			dl:       linkedlist.NewDoubly(),
			size:     2,
			capacity: 2,
			storage:  make(map[string]*linkedlist.Node),
		}
		lru.Put("a", 1)
		lru.Put("b", 2)

		// Act
		lru.Put("c", 3)

		// Assert
		if _, ok := lru.storage["a"]; ok {
			t.Errorf("Scenario 2: Key 'a' should have been evicted from cache")
		}
		if _, ok := lru.storage["c"]; !ok {
			t.Errorf("Scenario 2: Key 'c' not added to the cache")
		}
	})

	// Scenario 3: Put new key in empty cache
	t.Run("put_new_key_in_empty_cache", func(t *testing.T) {
		// Arrange
		lru := &LRU{
			dl:       linkedlist.NewDoubly(),
			size:     0,
			capacity: 1,
			storage:  make(map[string]*linkedlist.Node),
		}

		// Act
		lru.Put("key", 123)

		// Assert
		if _, ok := lru.storage["key"]; !ok {
			t.Errorf("Scenario 3: Key 'key' not added to the cache")
		}
	})

	// Scenario 4: Put key with empty value
	t.Run("put_key_with_empty_value", func(t *testing.T) {
		// Arrange
		lru := &LRU{
			dl:       linkedlist.NewDoubly(),
			size:     0,
			capacity: 3,
			storage:  make(map[string]*linkedlist.Node),
		}

		// Act
		lru.Put("empty", "")

		// Assert
		if _, ok := lru.storage["empty"]; !ok {
			t.Errorf("Scenario 4: Key 'empty' not added to the cache")
		}
	})

	// Scenario 5: Put key with special characters
	t.Run("put_key_with_special_characters", func(t *testing.T) {
		// Arrange
		lru := &LRU{
			dl:       linkedlist.NewDoubly(),
			size:     0,
			capacity: 3,
			storage:  make(map[string]*linkedlist.Node),
		}

		// Act
		lru.Put("#$%", "special")

		// Assert
		if _, ok := lru.storage["#$%"]; !ok {
			t.Errorf("Scenario 5: Key '#$%' not added to the cache")
		}
	})

	// Scenario 6: Put key exceeding capacity
	t.Run("put_key_exceeding_capacity", func(t *testing.T) {
		// Arrange
		lru := &LRU{
			dl:       linkedlist.NewDoubly(),
			size:     2,
			capacity: 2,
			storage:  make(map[string]*linkedlist.Node),
		}
		lru.Put("x", 10)
		lru.Put("y", 20)

		// Act
		lru.Put("z", 30)
		lru.Put("w", 40)

		// Assert
		if len(lru.storage) != 2 {
			t.Errorf("Scenario 6: Cache should only contain keys 'z' and 'w'")
		}
		if _, ok := lru.storage["z"]; !ok || _, ok := lru.storage["w"]; !ok {
			t.Errorf("Scenario 6: Keys 'z' and 'w' not retained in the cache")
		}
	})
}
