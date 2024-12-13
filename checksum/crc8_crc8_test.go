package checksum_test

import (
	"testing"

	"github.com/TheAlgorithms/Go/checksum"
)

// TestCRC8 tests the CRC8 function with different scenarios
func TestCRC8(t *testing.T) {
	// Scenario 1: Calculate CRC8 Checksum for Normal Data
	t.Run("Calculate CRC8 Checksum for Normal Data", func(t *testing.T) {
		model := checksum.CRCModel{
			Poly:   0x07,
			Init:   0x00,
			RefIn:  false,
			RefOut: false,
			XorOut: 0x00,
			Name:   "CRC-8",
		}
		data := []byte{0x01, 0x02, 0x03}
		expectedChecksum := uint8(0xF4)
		actualChecksum := checksum.CRC8(data, model)

		if actualChecksum != expectedChecksum {
			t.Errorf("Scenario 1: Expected checksum %02X, got %02X", expectedChecksum, actualChecksum)
		}
		t.Log("Scenario 1: Successfully calculated CRC8 checksum for normal data")
	})

	// Scenario 2: Calculate CRC8 Checksum with Reflected Output
	t.Run("Calculate CRC8 Checksum with Reflected Output", func(t *testing.T) {
		model := checksum.CRCModel{
			Poly:   0x07,
			Init:   0x00,
			RefIn:  false,
			RefOut: true,
			XorOut: 0x00,
			Name:   "CRC-8",
		}
		data := []byte{0x01, 0x02, 0x03}
		expectedChecksum := uint8(0x2F)
		actualChecksum := checksum.CRC8(data, model)

		if actualChecksum != expectedChecksum {
			t.Errorf("Scenario 2: Expected checksum %02X, got %02X", expectedChecksum, actualChecksum)
		}
		t.Log("Scenario 2: Successfully calculated CRC8 checksum with reflected output")
	})

	// Scenario 3: Calculate CRC8 Checksum with Zero-Length Data
	t.Run("Calculate CRC8 Checksum with Zero-Length Data", func(t *testing.T) {
		model := checksum.CRCModel{
			Poly:   0x07,
			Init:   0xAA,
			RefIn:  false,
			RefOut: false,
			XorOut: 0x00,
			Name:   "CRC-8",
		}
		data := []byte{}
		expectedChecksum := uint8(0xAA)
		actualChecksum := checksum.CRC8(data, model)

		if actualChecksum != expectedChecksum {
			t.Errorf("Scenario 3: Expected checksum %02X, got %02X", expectedChecksum, actualChecksum)
		}
		t.Log("Scenario 3: Successfully handled zero-length data for CRC8 calculation")
	})

	// Scenario 4: Calculate CRC8 Checksum with Maximum Data Length
	t.Run("Calculate CRC8 Checksum with Maximum Data Length", func(t *testing.T) {
		model := checksum.CRCModel{
			Poly:   0x07,
			Init:   0x00,
			RefIn:  false,
			RefOut: false,
			XorOut: 0xFF,
			Name:   "CRC-8",
		}
		data := make([]byte, 256) // Maximum supported length
		expectedChecksum := uint8(0x00)
		actualChecksum := checksum.CRC8(data, model)

		if actualChecksum != expectedChecksum {
			t.Errorf("Scenario 4: Expected checksum %02X, got %02X", expectedChecksum, actualChecksum)
		}
		t.Log("Scenario 4: Successfully calculated CRC8 checksum with maximum data length")
	})

	// Scenario 5: Calculate CRC8 Checksum with Invalid CRCModel
	t.Run("Calculate CRC8 Checksum with Invalid CRCModel", func(t *testing.T) {
		model := checksum.CRCModel{
			Poly:   0xFF, // Invalid polynomial
			Init:   0x00,
			RefIn:  false,
			RefOut: false,
			XorOut: 0x00,
			Name:   "CRC-8",
		}
		data := []byte{0x01, 0x02, 0x03}
		expectedChecksum := uint8(0x00) // As the model is invalid, expect 0x00
		actualChecksum := checksum.CRC8(data, model)

		if actualChecksum != expectedChecksum {
			t.Errorf("Scenario 5: Expected checksum %02X, got %02X", expectedChecksum, actualChecksum)
		}
		t.Log("Scenario 5: Successfully handled invalid CRCModel for CRC8 calculation")
	})
}
