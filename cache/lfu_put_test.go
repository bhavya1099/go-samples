package cache_test

import (
	"container/list"
	"math"
	"testing"

	"github.com/TheAlgorithms/Go/cache"
)

func TestPut(t *testing.T) {
	tests := []struct {
		name         string
		cacheSize    int
		existingKeys map[string]interface{}
		newKey       string
		newValue     interface{}
		expectedSize int
	}{
		{
			name:      "Put existing key with updated value",
			cacheSize: 3,
			existingKeys: map[string]interface{}{
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
			},
			newKey:       "key2",
			newValue:     "updatedValue2",
			expectedSize: 3,
		},
		{
			name:      "Put new key within capacity",
			cacheSize: 3,
			existingKeys: map[string]interface{}{
				"key1": "value1",
				"key2": "value2",
			},
			newKey:       "key3",
			newValue:     "value3",
			expectedSize: 3,
		},
		{
			name:      "Put new key exceeding capacity",
			cacheSize: 2,
			existingKeys: map[string]interface{}{
				"key1": "value1",
				"key2": "value2",
			},
			newKey:       "key3",
			newValue:     "value3",
			expectedSize: 2,
		},
		{
			name:      "Put key with empty value",
			cacheSize: 3,
			existingKeys: map[string]interface{}{
				"key1": "value1",
				"key2": "value2",
			},
			newKey:       "key3",
			newValue:     "",
			expectedSize: 3,
		},
		{
			name:      "Put key with special characters",
			cacheSize: 3,
			existingKeys: map[string]interface{}{
				"key1": "value1",
				"key2": "value2",
			},
			newKey:       "#$%key",
			newValue:     "specialValue",
			expectedSize: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			lfu := cache.LFU{
				len:     0,
				cap:     tt.cacheSize,
				minFreq: math.MaxInt32,
				itemMap: make(map[string]*list.Element),
				freqMap: make(map[int]*list.List),
			}

			for key, value := range tt.existingKeys {
				lfu.Put(key, value)
			}

			// Act
			lfu.Put(tt.newKey, tt.newValue)

			// Assert
			if lfu.len != tt.expectedSize {
				t.Errorf("Expected cache size: %d, got: %d", tt.expectedSize, lfu.len)
			}

			if tt.newValue != "" {
				element := lfu.itemMap[tt.newKey]
				if element == nil {
					t.Errorf("New key not found in cache")
				} else {
					item := element.Value.(cache.Item)
					if item.Value != tt.newValue {
						t.Errorf("Expected value: %v, got: %v", tt.newValue, item.Value)
					}
				}
			}
		})
	}
}
