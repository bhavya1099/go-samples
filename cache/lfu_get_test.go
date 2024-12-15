package cache

import (
	"container/list"
	"math"
	"testing"
)

func TestLFUGet(t *testing.T) {
	// Scenario 1: Existing Key - Normal Operation
	t.Run("Scenario 1: Existing Key - Normal Operation", func(t *testing.T) {
		cache := LFU{
			len:     0,
			cap:     10,
			minFreq: math.MaxInt32,
			itemMap: make(map[string]*list.Element),
			freqMap: make(map[int]*list.List),
		}

		key := "existing_key"
		value := "existing_value"
		cache.itemMap[key] = cache.addFreqNode(1, key, value) // Simulate adding an item to the cache with frequency 1

		t.Log("Testing with existing key")
		result := cache.Get(key)
		if result != value {
			t.Errorf("Scenario 1: Existing Key - Normal Operation failed. Expected value: %s, Got: %s", value, result)
		}
	})

	// Scenario 2: Non-Existing Key - Normal Operation
	t.Run("Scenario 2: Non-Existing Key - Normal Operation", func(t *testing.T) {
		cache := LFU{
			len:     0,
			cap:     10,
			minFreq: math.MaxInt32,
			itemMap: make(map[string]*list.Element),
			freqMap: make(map[int]*list.List),
		}

		key := "non_existing_key"

		t.Log("Testing with non-existing key")
		result := cache.Get(key)
		if result != nil {
			t.Errorf("Scenario 2: Non-Existing Key - Normal Operation failed. Expected nil, Got: %v", result)
		}
	})

	// Scenario 3: Key with Minimum Frequency - Edge Case
	t.Run("Scenario 3: Key with Minimum Frequency - Edge Case", func(t *testing.T) {
		cache := LFU{
			len:     0,
			cap:     10,
			minFreq: math.MaxInt32,
			itemMap: make(map[string]*list.Element),
			freqMap: make(map[int]*list.List),
		}

		key := "min_frequency_key"
		value := "min_frequency_value"
		cache.itemMap[key] = cache.addFreqNode(1, key, value) // Simulate adding an item to the cache with frequency 1

		t.Log("Testing with key having minimum frequency")
		result := cache.Get(key)
		if result != value {
			t.Errorf("Scenario 3: Key with Minimum Frequency - Edge Case failed. Expected value: %s, Got: %s", value, result)
		}
	})

	// Scenario 4: Large Cache Capacity - Edge Case
	t.Run("Scenario 4: Large Cache Capacity - Edge Case", func(t *testing.T) {
		cache := LFU{
			len:     0,
			cap:     2, // Setting a small capacity to reach the limit quickly
			minFreq: math.MaxInt32,
			itemMap: make(map[string]*list.Element),
			freqMap: make(map[int]*list.List),
		}

		key1, value1 := "key1", "value1"
		key2, value2 := "key2", "value2"
		key3, value3 := "key3", "value3"

		cache.itemMap[key1] = cache.addFreqNode(1, key1, value1)
		cache.itemMap[key2] = cache.addFreqNode(1, key2, value2)

		t.Log("Testing with large cache capacity")
		cache.Get(key1) // Increase frequency of key1
		cache.Get(key2) // Increase frequency of key2

		result := cache.Get(key1)
		if result != value1 {
			t.Errorf("Scenario 4: Large Cache Capacity - Edge Case failed. Expected value: %s, Got: %s", value1, result)
		}

		cache.itemMap[key3] = cache.addFreqNode(1, key3, value3) // Add a new key to reach capacity limit

		if cache.len != 2 {
			t.Errorf("Scenario 4: Large Cache Capacity - Edge Case failed. Cache size should be limited to 2.")
		}
	})

	// Scenario 5: Concurrent Access - Error Handling
	t.Run("Scenario 5: Concurrent Access - Error Handling", func(t *testing.T) {
		cache := LFU{
			len:     0,
			cap:     10,
			minFreq: math.MaxInt32,
			itemMap: make(map[string]*list.Element),
			freqMap: make(map[int]*list.List),
		}

		keys := []string{"key1", "key2", "key3"}

		for _, key := range keys {
			cache.itemMap[key] = cache.addFreqNode(1, key, fmt.Sprintf("value_%s", key))
		}

		var results []interface{}
		done := make(chan struct{})

		for _, key := range keys {
			go func(k string) {
				results = append(results, cache.Get(k))
				done <- struct{}{}
			}(key)
		}

		for range keys {
			<-done
		}

		for i, key := range keys {
			if results[i] != fmt.Sprintf("value_%s", key) {
				t.Errorf("Scenario 5: Concurrent Access - Error Handling failed. Incorrect value retrieved for key: %s", key)
			}
		}
	})
}
