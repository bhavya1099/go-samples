/*
  package cache_test
*/

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"cache"
	"math"
)

func TestinitItem(t *testing.T) {
	type testItem struct {
		key   string
		value interface{}
		freq  int
	}

	tests := []struct {
		name     string
		input    testItem
		expected cache.Item
	}{
		{
			name: "Scenario 1: Initialize Item with Positive Frequency",
			input: testItem{
				key:   "test_key",
				value: 123,
				freq:  5,
			},
			expected: cache.Item{
				key:   "test_key",
				value: 123,
				freq:  5,
			},
		},
		{
			name: "Scenario 2: Initialize Item with Zero Frequency",
			input: testItem{
				key:   "zero_freq",
				value: "initial",
				freq:  0,
			},
			expected: cache.Item{
				key:   "zero_freq",
				value: "initial",
				freq:  0,
			},
		},
		{
			name: "Scenario 3: Initialize Item with Negative Frequency",
			input: testItem{
				key:   "negative_freq",
				value: -99,
				freq:  -3,
			},
			expected: cache.Item{
				key:   "negative_freq",
				value: -99,
				freq:  -3,
			},
		},
		{
			name: "Scenario 4: Initialize Item with Large Value",
			input: testItem{
				key:   "large_value",
				value: math.MaxInt64,
				freq:  1,
			},
			expected: cache.Item{
				key:   "large_value",
				value: math.MaxInt64,
				freq:  1,
			},
		},
		{
			name: "Scenario 5: Initialize Item with Empty Key and Value",
			input: testItem{
				key:   "",
				value: nil,
				freq:  2,
			},
			expected: cache.Item{
				key:   "",
				value: nil,
				freq:  2,
			},
		},
		{
			name: "Scenario 6: Initialize Item with Special Characters Key",
			input: testItem{
				key:   "!@#$%^&*",
				value: "special",
				freq:  4,
			},
			expected: cache.Item{
				key:   "!@#$%^&*",
				value: "special",
				freq:  4,
			},
		},
		{
			name: "Scenario 7: Initialize Item with Negative Value",
			input: testItem{
				key:   "neg_value",
				value: -50,
				freq:  1,
			},
			expected: cache.Item{
				key:   "neg_value",
				value: -50,
				freq:  1,
			},
		},
		{
			name: "Scenario 8: Initialize Item with Maximum Frequency",
			input: testItem{
				key:   "max_freq",
				value: "max",
				freq:  math.MaxInt32,
			},
			expected: cache.Item{
				key:   "max_freq",
				value: "max",
				freq:  math.MaxInt32,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := cache.InitItem(tt.input.key, tt.input.value, tt.input.freq)
			assert.Equal(t, tt.expected, result, "Test case %s failed: Expected %v, got %v", tt.name, tt.expected, result)
		})
	}
}
