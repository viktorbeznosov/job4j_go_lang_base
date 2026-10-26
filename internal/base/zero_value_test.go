package base

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBagWithoutItems(t *testing.T) {
    t.Run("TestBagWithoutItems", func(t *testing.T) {
        t.Parallel()

        bag := new(Bag)

        assert.Equal(t, bag.Len(), 0)
        assert.False(t, bag.Contains("anything"))
    })
}

func TestBagWithItems(t *testing.T) {
    t.Run("TestBagWithItems", func(t *testing.T) {
        t.Parallel()

        bag := new(Bag)
        foo := "foo"
        bar := "bar"
        baz := "baz"
        bag.Add(foo)
        bag.Add(bar)

        assert.Equal(t, bag.Len(), 2)
        assert.True(t, bag.Contains(foo))
        assert.True(t, bag.Contains(bar))
        assert.False(t, bag.Contains(baz))
    })
}

func TestWhenAddSameElements(t *testing.T) {
    t.Run("TestWhenAddSameElements", func(t *testing.T) {
        t.Parallel()

        bag := Bag{}
        bag.Add("item")
        bag.Add("item")
        bag.Add("item")

        assert.Equal(t, bag.Len(), 1)
    })
}

func TestWhenEmptyString(t *testing.T) {
    t.Run("TestWhenAddSameElements", func(t *testing.T) {
        t.Parallel()

        bag := Bag{}
        bag.Add("")

        assert.Equal(t, bag.Len(), 1)
        assert.True(t, bag.Contains(""))
    })
}
