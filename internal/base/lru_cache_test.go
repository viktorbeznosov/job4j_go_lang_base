package base_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/viktorbeznosov/job4j_go_lang_base/internal/base"
)

func Test_Create_Lru_Cache(t *testing.T) {
	lruCache := base.NewLruCache(3)

	lruCache.Put("1", "Value1")
	lruCache.Put("2", "Value2")
	lruCache.Put("3", "Value3")

	assert.Equal(t, lruCache.Length(), 3)
}

func Test_Overflow_Сontrol(t *testing.T) {
	lruCache := base.NewLruCache(3)

	lruCache.Put("1", "Value1")
	lruCache.Put("2", "Value2")
	lruCache.Put("3", "Value3")
	lruCache.Put("4", "Value4")

	assert.Equal(t, lruCache.Length(), 3)
	assert.Equal(t, lruCache.Head.Value, "Value4")
	assert.Equal(t, lruCache.Tail.Value, "Value2")
}

func Test_Get_Elements(t *testing.T) {
	lruCache := base.NewLruCache(3)

	lruCache.Put("1", "Value1")
	lruCache.Put("2", "Value2")
	lruCache.Put("3", "Value3")
	lruCache.Put("2", "Another Value")

	foundElement := lruCache.Get("2")

	assert.Equal(t, lruCache.Length(), 3)
	assert.NotEmpty(t, foundElement)
	assert.Equal(t, *foundElement, "Another Value")
	assert.Equal(t, lruCache.Head.Key, "2")
}
