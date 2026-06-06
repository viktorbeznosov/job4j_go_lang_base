package main

import (
	"fmt"

	"github.com/viktorbeznosov/job4j_go_lang_base/internal/base"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println(base.Add(1, 3))
	fmt.Println(base.Max(3, 2))
	fmt.Println(base.Count(5))

	sl := []int{1, 1, 2, 3, 1}
	fmt.Printf("%t\n", base.Mono(sl))
}
