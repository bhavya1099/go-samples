package checksum

import "testing"

func TestLuhn(t *testing.T) {
	testCases := []struct {
		name     string
		input    []byte
		expected bool
	}{
		{
			name:     "Valid Luhn Checksum for Even Length Input",
			input:    []byte{5, 4, 9, 3, 1, 4, 8, 8, 8, 8, 8, 8, 8, 8},
			expected: true,
		},
		{
			name:     "Valid Luhn Checksum for Odd Length Input",
			input:    []byte{5, 4, 9, 3, 1, 4, 8, 8, 8, 8, 8, 8, 8},
			expected: true,
		},
		{
			name:     "Invalid Luhn Checksum for Even Length Input",
			input:    []byte{5, 4, 9, 3, 1, 4, 8, 8, 8, 8, 8, 8, 8, 7},
			expected: false,
		},
		{
			name:     "Invalid Luhn Checksum for Odd Length Input",
			input:    []byte{5, 4, 9, 3, 1, 4, 8, 8, 8, 8, 8, 8, 7},
			expected: false,
		},
		{
			name:     "Empty Input",
			input:    []byte{},
			expected: false,
		},
		{
			name:     "Large Input",
			input:    generateLargeInput(),
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if result := Luhn(tc.input); result != tc.expected {
				t.Errorf("Test Case: %s failed. Expected: %v, Got: %v", tc.name, tc.expected, result)
			}
		})
	}
}
func generateLargeInput() []byte {

	return []byte{}
}
