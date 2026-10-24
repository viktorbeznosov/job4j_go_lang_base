package base

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCelsius(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		raw      string
		expected int16
		wantErr  bool
	}{
		{
			name:     "zero",
			raw:      "0",
			expected: 0,
			wantErr:  false,
		},
		{
			name:     "min int16 boundary",
			raw:      "-32768",
			expected: -32768,
			wantErr:  false,
		},
		{
			name:     "max int16 boundary",
			raw:      "32767",
			expected: 32767,
			wantErr:  false,
		},
		{
			name:     "negative value",
			raw:      "-40",
			expected: -40,
			wantErr:  false,
		},
		{
			name:     "positive value",
			raw:      "25",
			expected: 25,
			wantErr:  false,
		},
		{
			name:     "overflow above max int16",
			raw:      "32768",
			expected: 0,
			wantErr:  true,
		},
		{
			name:     "overflow below min int16",
			raw:      "-32769",
			expected: 0,
			wantErr:  true,
		},
		{
			name:     "invalid string",
			raw:      "abc",
			expected: 0,
			wantErr:  true,
		},
		{
			name:     "empty string",
			raw:      "",
			expected: 0,
			wantErr:  true,
		},
		{
			name:     "float string",
			raw:      "12.5",
			expected: 0,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result, err := ParseCelsius(tt.raw)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Equal(t, int16(0), result)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestToFahrenheit(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		celsius  int16
		expected float64
	}{
		{
			name:     "zero",
			celsius:  0,
			expected: 32,
		},
		{
			name:     "min int16 boundary",
			celsius:  -32768,
			expected: -58950,
		},
		{
			name:     "max int16 boundary",
			celsius:  32767,
			expected: 59012,
		},
		{
			name:     "negative value",
			celsius:  -40,
			expected: -40,
		},
		{
			name:     "positive value",
			celsius:  25,
			expected: 77,
		},
		{
			name:     "truncation case",
			celsius:  11,
			expected: 51,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := ToFahrenheit(tt.celsius)
			assert.Equal(t, tt.expected, result)
		})
	}
}
