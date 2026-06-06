package base_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/viktorbeznosov/job4j_go_lang_base/internal/base"
)

func Test_Add(t *testing.T) {
	rsl := base.Add(1, 2)
	expected := 3

	assert.Equal(t, expected, rsl, "Add(1,2) should equal 3")
}

func Test_Max(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"first greater", 5, 3, 5},
		{"second greater", 2, 8, 8},
		{"equal", 4, 4, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := base.Max(tt.a, tt.b)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func Test_Count(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected int
	}{
		{"n=5", 5, 15},
		{"n=0", 0, 0},
		{"n=1", 1, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := base.Count(tt.n)
			assert.Equal(t, tt.expected, result)
		})
	}
}
