package main

import (
	"fmt"

	"github.com/viktorbeznosov/job4j_go_lang_base/internal/base"
)

func main() {
	lruCache := base.NewLruCache(5)

	lruCache.Put("1", "Value1")
	lruCache.Put("2", "Value2")
	lruCache.Put("3", "Value3")
	lruCache.Put("2", "QWERTY")
	lruCache.Put("4", "Value4")
	lruCache.Put("5", "Value5")
	fmt.Println(lruCache.Length())

	lruCache.Get("3")

	current := lruCache.Head
	for current != nil {
		fmt.Println(current)
		current = current.Prev
	}
}
