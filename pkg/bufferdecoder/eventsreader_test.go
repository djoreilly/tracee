package bufferdecoder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadStringVarFromBuff(t *testing.T) {
	tests := []struct {
		name           string
		buffer         []byte
		max            int
		expected       string
		expectedCursor int
		expectError    bool
	}{
		{
			name:           "Null terminated string in larger buffer",
			buffer:         []byte{'H', 'e', 'l', 'l', 'o', 0, 'W', 'o', 'r', 'l', 'd'},
			max:            6,
			expected:       "Hello",
			expectedCursor: 6,
			expectError:    false,
		},
		{
			name:           "Buffer with same length as max without null terminator",
			buffer:         []byte{'H', 'e', 'l', 'l', 'o'},
			max:            5,
			expected:       "Hell",
			expectedCursor: 5,
			expectError:    false,
		},
		{
			name:           "Buffer longer than max length without null terminator",
			buffer:         []byte{'H', 'e', 'l', 'l', 'o', 'W', 'o', 'r', 'l', 'd'},
			max:            5,
			expected:       "Hell",
			expectedCursor: 5,
			expectError:    false,
		},
		{
			name:           "Zero max length",
			buffer:         []byte{'H', 'e', 'l', 'l', 'o', 0, 'W', 'o', 'r', 'l', 'd'},
			max:            0,
			expected:       "",
			expectedCursor: 0,
			expectError:    false,
		},
		{
			name:           "Buffer started with null terminator",
			buffer:         []byte{0, 'N', 'u', 'l', 'l', 0, 'W', 'o', 'r', 'l', 'd'},
			max:            6,
			expected:       "",
			expectedCursor: 6,
			expectError:    false,
		},
		{
			name:           "Empty buffer",
			buffer:         []byte{},
			max:            5,
			expected:       "",
			expectedCursor: 0,
			expectError:    true,
		},
		{
			name:           "Buffer too short",
			buffer:         []byte{'H'},
			max:            5,
			expected:       "H",
			expectedCursor: 0,
			expectError:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decoder := New(tt.buffer)
			actual, err := readStringVarFromBuff(decoder, tt.max)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, actual)
				assert.Equal(t, tt.expectedCursor, decoder.ReadAmountBytes())
			}
		})
	}
}

func TestPrintUint32IP(t *testing.T) {
	t.Parallel()

	var input uint32 = 3232238339
	ip := PrintUint32IP(input)

	expectedIP := "192.168.11.3"
	assert.Equal(t, expectedIP, ip)
}

func TestPrint16BytesSliceIP(t *testing.T) {
	t.Parallel()

	input := []byte{32, 1, 13, 184, 133, 163, 0, 0, 0, 0, 138, 46, 3, 112, 115, 52}
	ip := Print16BytesSliceIP(input)

	expectedIP := "2001:db8:85a3::8a2e:370:7334"
	assert.Equal(t, expectedIP, ip)
}
