package checksum

import (
	"testing"
	"your/module/checksum"
)

func TestgetTable(t *testing.T) {
	testCases := []struct {
		name          string
		model         checksum.CRCModel
		expectedTable []uint8
	}{
		{
			name: "Standard CRC Model Test",
			model: checksum.CRCModel{
				Poly:   0x1D,
				Init:   0x00,
				RefIn:  false,
				RefOut: false,
				XorOut: 0x00,
				Name:   "CRC-8",
			},
			expectedTable: []uint8{
				0x00, 0x1D, 0x3A, 0x27, 0x74, 0x69, 0x4E, 0x53,
				0xE8, 0xF5, 0xD2, 0xCF, 0x9C, 0x81, 0xA6, 0xBB,
			},
		},
		{
			name: "Custom CRC Model Test",
			model: checksum.CRCModel{
				Poly:   0x8C,
				Init:   0xFF,
				RefIn:  true,
				RefOut: true,
				XorOut: 0x55,
				Name:   "Custom-CRC",
			},
			expectedTable: []uint8{
				0xFF, 0x73, 0xE6, 0x95, 0x2B, 0xC7, 0x4E, 0x24,
				0x2A, 0x2C, 0x53, 0x52, 0x0F, 0x3E, 0x2D, 0x42,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			table := checksum.GetTable(tc.model)

			if len(table) != 256 {
				t.Errorf("Expected table length to be 256, got %d", len(table))
			}

			for i, val := range tc.expectedTable {
				if table[i] != val {
					t.Errorf("Table value mismatch at index %d. Expected: %d, Got: %d", i, val, table[i])
				}
			}
		})
	}
}
