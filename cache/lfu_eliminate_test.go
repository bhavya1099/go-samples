package cache

import (
	"container/list"
	"math"
	"testing"
)

func TestLFUeliminate(t *testing.T) {
	// Scenario 1
	t.Run("Test LFU eliminate function with a valid item to remove", func(t *testing.T) {
		lfu := LFU{
			len:     0,
			cap:     5,
			minFreq: math.MaxInt32,
			itemMap: make(map[string]*list.Element),
			freqMap: make(map[int]*list.List),
		}

		// Add an item with freq 1
		item1 := item{
			key:   "key1",
			value: "value1",
			freq:  1,
		}
		l := list.New()
		l.PushFront(item1)
		lfu.itemMap["key1"] = l.Front()
		lfu.freqMap[1] = l

		lfu.eliminate()

		// Check if item with freq 1 is removed
		if _, ok := lfu.itemMap["key1"]; ok {
			t.Error("Failed to remove item with minimum frequency from LFU cache")
		}
	})

	// Scenario 2
	t.Run("Test LFU eliminate function with an empty cache", func(t *testing.T) {
		lfu := LFU{
			len:     0,
			cap:     5,
			minFreq: math.MaxInt32,
			itemMap: make(map[string]*list.Element),
			freqMap: make(map[int]*list.List),
		}

		// Call eliminate on empty cache
		lfu.eliminate()

		// No error should occur
	})

	// Scenario 3
	t.Run("Test LFU eliminate function with multiple items at the same frequency", func(t *testing.T) {
		lfu := LFU{
			len:     0,
			cap:     5,
			minFreq: math.MaxInt32,
			itemMap: make(map[string]*list.Element),
			freqMap: make(map[int]*list.List),
		}

		// Add items with freq 2
		item1 := item{
			key:   "key1",
			value: "value1",
			freq:  2,
		}
		item2 := item{
			key:   "key2",
			value: "value2",
			freq:  2,
		}
		l := list.New()
		l.PushFront(item1)
		lfu.itemMap["key1"] = l.Front()
		l.PushFront(item2)
		lfu.itemMap["key2"] = l.Front()
		lfu.freqMap[2] = l

		lfu.eliminate()

		// Check if the correct item is removed
		if _, ok := lfu.itemMap["key1"]; ok {
			t.Error("Failed to remove correct item based on LFU eviction policy")
		}
	})

	// Scenario 4
	t.Run("Test LFU eliminate function with a full cache", func(t *testing.T) {
		lfu := LFU{
			len:     0,
			cap:     2,
			minFreq: math.MaxInt32,
			itemMap: make(map[string]*list.Element),
			freqMap: make(map[int]*list.List),
		}

		// Fill cache with items
		item1 := item{
			key:   "key1",
			value: "value1",
			freq:  1,
		}
		item2 := item{
			key:   "key2",
			value: "value2",
			freq:  2,
		}
		l1 := list.New()
		l1.PushFront(item1)
		lfu.itemMap["key1"] = l1.Front()
		lfu.freqMap[1] = l1

		l2 := list.New()
		l2.PushFront(item2)
		lfu.itemMap["key2"] = l2.Front()
		lfu.freqMap[2] = l2

		lfu.eliminate()

		// Check if item with freq 1 is removed
		if _, ok := lfu.itemMap["key1"]; ok {
			t.Error("Failed to remove item with minimum frequency from full LFU cache")
		}
	})

	// Scenario 5
	t.Run("Test LFU eliminate function with a single item in the cache", func(t *testing.T) {
		lfu := LFU{
			len:     0,
			cap:     1,
			minFreq: math.MaxInt32,
			itemMap: make(map[string]*list.Element),
			freqMap: make(map[int]*list.List),
		}

		// Add a single item
		item1 := item{
			key:   "key1",
			value: "value1",
			freq:  1,
		}
		l := list.New()
		l.PushFront(item1)
		lfu.itemMap["key1"] = l.Front()
		lfu.freqMap[1] = l

		lfu.eliminate()

		// Check if item is removed
		if _, ok := lfu.itemMap["key1"]; ok {
			t.Error("Failed to remove single item from LFU cache")
		}
	})
}
