package checksum

import (
	"testing"
	"math/bits"
)

func TestaddBytes(t *testing.T) {

	t.Run("Scenario 1: Test normal operation without bit reversal", func(t *testing.T) {

		model := CRCModel{Poly: 0x07, Init: 0x00, RefIn: false, RefOut: false, XorOut: 0x00, Name: "CRC-8"}
		data := []byte{0x01, 0x02, 0x03}
		crcResult := uint8(0x00)
		table := generateTable(model)

		result := addBytes(data, model, crcResult, table)

		expected := uint8(0xD4)
		if result != expected {
			t.Errorf("Scenario 1: Expected %d, but got %d", expected, result)
		}
	})

	t.Run("Scenario 2: Test normal operation with bit reversal", func(t *testing.T) {

		model := CRCModel{Poly: 0x07, Init: 0x00, RefIn: true, RefOut: false, XorOut: 0x00, Name: "CRC-8"}
		data := []byte{0x01, 0x02, 0x03}
		crcResult := uint8(0x00)
		table := generateTable(model)

		result := addBytes(data, model, crcResult, table)

		expected := uint8(0x2B)
		if result != expected {
			t.Errorf("Scenario 2: Expected %d, but got %d", expected, result)
		}
	})
}
func generateTable(model CRCModel) []uint8 {
	table := make([]uint8, 256)
	for i := range table {
		crc := uint8(i)
		for j := 0; j < 8; j++ {
			if crc&0x80 != 0 {
				crc = (crc << 1) ^ model.Poly
			} else {
				crc <<= 1
			}
		}
		if model.RefOut {
			table[i] = bits.Reverse8(crc)
		} else {
			table[i] = crc
		}
	}
	return table
}
