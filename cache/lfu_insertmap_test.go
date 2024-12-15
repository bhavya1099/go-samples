package cache_test

import (
	"container/list"
	"testing"

	"github.com/path/to/your/package/cache" // TODO: Update the import path
)

func TestinsertMap(t *testing.T) {
	// Scenario 1: Insert item into an empty frequency map
	t.Run("Insert into empty frequency map", func(t *testing.T) {
		// Arrange
		lfu := cache.LFU{
			FreqMap: make(map[int]*list.List),
			ItemMap: make(map[string]*list.Element),
		}
		item := cache.Item{
			Key:   "key1",
			Value: "value1",
			Freq:  1,
		}

		// Act
		lfu.InsertMap(item)

		// Assert
		if _, ok := lfu.FreqMap[item.Freq]; !ok {
			t.Error("Frequency map not updated")
		}
		if _, ok := lfu.ItemMap[item.Key]; !ok {
			t.Error("Item not added to item map")
		}
	})

	// Scenario 2: Insert item with existing frequency
	t.Run("Insert with existing frequency", func(t *testing.T) {
		// Arrange
		lfu := cache.LFU{
			FreqMap: map[int]*list.List{
				1: list.New(),
			},
			ItemMap: make(map[string]*list.Element),
		}
		item := cache.Item{
			Key:   "key2",
			Value: "value2",
			Freq:  1,
		}

		// Act
		lfu.InsertMap(item)

		// Assert
		if lfu.FreqMap[item.Freq].Len() != 1 {
			t.Error("Item not added to frequency list")
		}
		if _, ok := lfu.ItemMap[item.Key]; !ok {
			t.Error("Item not added to item map")
		}
	})

	// Scenario 3: Insert item with new frequency
	t.Run("Insert with new frequency", func(t *testing.T) {
		// Arrange
		lfu := cache.LFU{
			FreqMap: map[int]*list.List{
				1: list.New(),
			},
			ItemMap: make(map[string]*list.Element),
		}
		item := cache.Item{
			Key:   "key3",
			Value: "value3",
			Freq:  2,
		}

		// Act
		lfu.InsertMap(item)

		// Assert
		if _, ok := lfu.FreqMap[item.Freq]; !ok {
			t.Error("New frequency list not created")
		}
		if lfu.FreqMap[item.Freq].Len() != 1 {
			t.Error("Item not added to new frequency list")
		}
		if _, ok := lfu.ItemMap[item.Key]; !ok {
			t.Error("Item not added to item map")
		}
	})

	// Scenario 4: Insert item with duplicate key
	t.Run("Insert with duplicate key", func(t *testing.T) {
		// Arrange
		lfu := cache.LFU{
			FreqMap: map[int]*list.List{
				1: list.New(),
			},
			ItemMap: map[string]*list.Element{
				"key4": nil, // Existing item with key "key4"
			},
		}
		item := cache.Item{
			Key:   "key4",
			Value: "value4",
			Freq:  1,
		}

		// Act
		lfu.InsertMap(item)

		// Assert
		if lfu.ItemMap[item.Key] == nil {
			t.Error("Existing item with same key overridden")
		}
	})

	// Scenario 5: Insert item with empty key
	t.Run("Insert with empty key", func(t *testing.T) {
		// Arrange
		lfu := cache.LFU{
			FreqMap: make(map[int]*list.List),
			ItemMap: make(map[string]*list.Element),
		}
		item := cache.Item{
			Key:   "", // Empty key
			Value: "value5",
			Freq:  1,
		}

		// Act
		lfu.InsertMap(item)

		// Assert
		if _, ok := lfu.ItemMap[item.Key]; ok {
			t.Error("Item with empty key added to cache")
		}
	})
}
