package cache

import (
	"container/list"
	"math"
)

func TestincreaseFreq(t *testing.T) {
	// Scenario 1: Increase frequency for an element with existing frequency in the LFU cache
	t.Run("Increase frequency for existing element", func(t *testing.T) {
		// Arrange
		lfu := LFU{
			len:     0,
			cap:     10,
			minFreq: math.MaxInt64,
			itemMap: make(map[string]*list.Element),
			freqMap: make(map[int]*list.List),
		}
		key := "test_key"
		item := item{key: key, freq: 1}
		e := lfu.insertMap(item)

		// Act
		lfu.increaseFreq(e)

		// Assert
		if e.Value.(item).freq != 2 {
			t.Errorf("Expected frequency to be 2, got %d", e.Value.(item).freq)
		}
	})

	// Scenario 2: Increase frequency for the element with the minimum frequency in the LFU cache
	t.Run("Increase frequency for element with minimum frequency", func(t *testing.T) {
		// Arrange
		lfu := LFU{
			len:     0,
			cap:     10,
			minFreq: 1,
			itemMap: make(map[string]*list.Element),
			freqMap: make(map[int]*list.List),
		}
		key := "test_key"
		item := item{key: key, freq: 1}
		e := lfu.insertMap(item)

		// Act
		lfu.increaseFreq(e)

		// Assert
		if lfu.minFreq != 2 {
			t.Errorf("Expected minFreq to be 2, got %d", lfu.minFreq)
		}
	})

	// Scenario 3: Increase frequency for the only element in the LFU cache
	t.Run("Increase frequency for the only element", func(t *testing.T) {
		// Arrange
		lfu := LFU{
			len:     0,
			cap:     10,
			minFreq: math.MaxInt64,
			itemMap: make(map[string]*list.Element),
			freqMap: make(map[int]*list.List),
		}
		key := "test_key"
		item := item{key: key, freq: 1}
		e := lfu.insertMap(item)

		// Act
		lfu.increaseFreq(e)

		// Assert
		if e.Value.(item).freq != 2 {
			t.Errorf("Expected frequency to be 2, got %d", e.Value.(item).freq)
		}
	})

	// Scenario 4: Increase frequency for a non-existent element in the LFU cache
	t.Run("Increase frequency for non-existent element", func(t *testing.T) {
		// Arrange
		lfu := LFU{
			len:     0,
			cap:     10,
			minFreq: math.MaxInt64,
			itemMap: make(map[string]*list.Element),
			freqMap: make(map[int]*list.List),
		}
		key := "test_key"
		item := item{key: key, freq: 1}
		e := lfu.insertMap(item)

		// Act
		nonExistentKey := "non_existent_key"
		nonExistentItem := item{key: nonExistentKey, freq: 1}
		nonExistentE := lfu.insertMap(nonExistentItem)
		lfu.freqMap[1].Remove(nonExistentE) // Simulating non-existence

		// Assert
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected a panic for increasing frequency of a non-existent element")
			}
		}()
		lfu.increaseFreq(nonExistentE)
	})

	// Scenario 5: Increase frequency for an element with the maximum frequency in the LFU cache
	t.Run("Increase frequency for element with maximum frequency", func(t *testing.T) {
		// Arrange
		lfu := LFU{
			len:     0,
			cap:     10,
			minFreq: math.MaxInt64,
			itemMap: make(map[string]*list.Element),
			freqMap: make(map[int]*list.List),
		}
		key := "test_key"
		item := item{key: key, freq: math.MaxInt64}
		e := lfu.insertMap(item)

		// Act
		lfu.increaseFreq(e)

		// Assert
		if e.Value.(item).freq != math.MaxInt64 {
			t.Errorf("Expected frequency to remain unchanged at %d, got %d", math.MaxInt64, e.Value.(item).freq)
		}
	})
}
