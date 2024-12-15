package cache_test

import (
	"testing"
	"github.com/TheAlgorithms/Go/structure/linkedlist"
	"github.com/your-package-path/cache" // TODO: Update with the correct package path
)

func TestNewLRU(t *testing.T) {
	testCases := []struct {
		name     string
		capacity int
	}{
		{
			name:     "Capacity greater than zero",
			capacity: 10,
		},
		{
			name:     "Capacity zero",
			capacity: 0,
		},
		{
			name:     "Negative capacity",
			capacity: -5,
		},
		{
			name:     "Large capacity",
			capacity: 1000,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Log("Running test for:", tc.name)

			lru := cache.NewLRU(tc.capacity)

			if tc.capacity < 0 {
				if lru.Size() != 0 || lru.Capacity() != 0 {
					t.Errorf("Expected size and capacity to be 0 for negative capacity, got: size=%d, capacity=%d", lru.Size(), lru.Capacity())
				}
			} else {
				if lru.Size() != 0 || lru.Capacity() != tc.capacity {
					t.Errorf("Expected size=0 and capacity=%d, got: size=%d, capacity=%d", tc.capacity, lru.Size(), lru.Capacity())
				}
			}

			if lru.GetDoublyLinkedList() == nil || lru.GetStorageMap() == nil {
				t.Error("Doubly linked list or storage map is nil")
			}

			if _, ok := lru.GetStorageMap()[""]; !ok {
				t.Error("Empty key not found in storage map")
			}

			if lru.GetDoublyLinkedList().Head() != nil || lru.GetDoublyLinkedList().Tail() != nil {
				t.Error("Doubly linked list is not empty")
			}
		})
	}
}
