package base

import "fmt"

type Counter struct {
    Value int
}

func (c *Counter) Print() {
    fmt.Println(c.Value)
}

func Run(withClosure bool) {
    c := Counter{Value: 123}
    if (withClosure) {
        defer func() {
            c.Print() //456
        }()
    } else {
        defer c.Print() //123
    }

    defer fmt.Println("First")
    defer fmt.Println("Middle")
    defer fmt.Println("Last")

    c.Value = 456
}

func Mixed() {
    value := 10
    defer func(saved int) {
        fmt.Println(saved, value)
    }(value)
    value = 20
} //10 20




