package cache_test

import (
	"container/list"
	"math"
	"testing"
	"cache" // Fixed import path
)

func TestNewLFU(t *testing.T) {
	tests := []struct {
		name      string
		capacity  int
		minFreq   int
		itemMap   map[string]*list.Element
		freqMap   map[int]*list.List
		expectErr bool
	}{
		{
			name:     "Test NewLFU Initialization with Zero Capacity",
			capacity: 0,
			minFreq:  math.MaxInt,
			itemMap:  make(map[string]*list.Element),
			freqMap:  make(map[int]*list.List),
		},
		{
			name:     "Test NewLFU Initialization with Positive Capacity",
			capacity: 10,
			minFreq:  math.MaxInt,
			itemMap:  make(map[string]*list.Element),
			freqMap:  make(map[int]*list.List),
		},
		{
			name:     "Test NewLFU Initialization with Large Capacity",
			capacity: 1000,
			minFreq:  math.MaxInt,
			itemMap:  make(map[string]*list.Element),
			freqMap:  make(map[int]*list.List),
		},
		{
			name:     "Test NewLFU Initialization with Negative Capacity",
			capacity: -5,
			minFreq:  math.MaxInt,
			itemMap:  make(map[string]*list.Element),
			freqMap:  make(map[int]*list.List),
		},
		{
			name:     "Test NewLFU Initialization with MaxInt Capacity",
			capacity: math.MaxInt,
			minFreq:  math.MaxInt,
			itemMap:  make(map[string]*list.Element),
			freqMap:  make(map[int]*list.List),
		},
		{
			name:     "Test NewLFU Initialization with Zero Capacity and Custom Initialization",
			capacity: 0,
			minFreq:  5,
			itemMap:  map[string]*list.Element{"test": nil},
			freqMap:  map[int]*list.List{5: list.New()},
		},
		{
			name:     "Test NewLFU Initialization with Positive Capacity and Custom Initialization",
			capacity: 10,
			minFreq:  3,
			itemMap:  map[string]*list.Element{"custom": nil},
			freqMap:  map[int]*list.List{3: list.New()},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lfu := cache.NewLFU(tt.capacity)

			if lfu.Capacity() != tt.capacity {
				t.Errorf("unexpected capacity, got: %d, want: %d", lfu.Capacity(), tt.capacity)
			}
			if lfu.MinFreq() != tt.minFreq {
				t.Errorf("unexpected minFreq, got: %d, want: %d", lfu.MinFreq(), tt.minFreq)
			}
			if len(lfu.ItemMap()) != len(tt.itemMap) {
				t.Errorf("unexpected itemMap length, got: %d, want: %d", len(lfu.ItemMap()), len(tt.itemMap))
			}
			if len(lfu.FreqMap()) != len(tt.freqMap) {
				t.Errorf("unexpected freqMap length, got: %d, want: %d", len(lfu.FreqMap()), len(tt.freqMap))
			}
		})
	}
}
