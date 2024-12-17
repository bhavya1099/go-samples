package checksum

import (
	"testing"
	"github.com/TheAlgorithms/Go/checksum"
)

func TestCRC8(t *testing.T) {

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

	t.Run("Calculate CRC8 Checksum with Maximum Data Length", func(t *testing.T) {
		model := checksum.CRCModel{
			Poly:   0x07,
			Init:   0x00,
			RefIn:  false,
			RefOut: false,
			XorOut: 0xFF,
			Name:   "CRC-8",
		}
		data := make([]byte, 256)
		expectedChecksum := uint8(0x00)
		actualChecksum := checksum.CRC8(data, model)

		if actualChecksum != expectedChecksum {
			t.Errorf("Scenario 4: Expected checksum %02X, got %02X", expectedChecksum, actualChecksum)
		}
		t.Log("Scenario 4: Successfully calculated CRC8 checksum with maximum data length")
	})

	t.Run("Calculate CRC8 Checksum with Invalid CRCModel", func(t *testing.T) {
		model := checksum.CRCModel{
			Poly:   0xFF,
			Init:   0x00,
			RefIn:  false,
			RefOut: false,
			XorOut: 0x00,
			Name:   "CRC-8",
		}
		data := []byte{0x01, 0x02, 0x03}
		expectedChecksum := uint8(0x00)
		actualChecksum := checksum.CRC8(data, model)

		if actualChecksum != expectedChecksum {
			t.Errorf("Scenario 5: Expected checksum %02X, got %02X", expectedChecksum, actualChecksum)
		}
		t.Log("Scenario 5: Successfully handled invalid CRCModel for CRC8 calculation")
	})
}
