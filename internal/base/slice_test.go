package base

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveValue(t *testing.T) {
	t.Parallel()

    tests := []struct {
        name     string
        delete   int
        raw      []int
        expected []int
    }{
        {
            name: "Test1",
            delete: 2,
            raw: []int{1, 2, 3},
            expected: []int{1, 3},
        },
        {
            name: "Test1",
            delete: 2,
            raw: []int{},
            expected: []int{},
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            t.Parallel()

            result := RemoveValue(tt.raw, tt.delete)
            assert.Equal(t, result, tt.expected)
        })
    }
}

func TestKeepPositive(t *testing.T) {
	t.Parallel()

    tests := []struct {
        name     string
        delete   int
        raw      []int
        expected []int
    }{
        {
            name: "Test1",
            raw: []int{1, -2, 3},
            expected: []int{1, 3},
        },
        {
            name: "Test1",
            raw: []int{1, 2, 3},
            expected: []int{1, 2, 3},
        },
        {
            name: "Test1",
            raw: []int{},
            expected: []int{},
        },
        {
            name: "Test1",
            raw: []int{-1, -2, -3},
            expected: []int{},
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            t.Parallel()

            result := KeepPositive(tt.raw)
            assert.Equal(t, result, tt.expected)
        })
    }
}


