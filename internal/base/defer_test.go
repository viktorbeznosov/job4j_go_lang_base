package base

import (
	"reflect"
	"testing"
)

func TestRun_DeferOrder(t *testing.T) {
    t.Parallel()
    var order []string

    func() {
        defer func() { order = append(order, "first") }()
        defer func() { order = append(order, "middle") }()
        defer func() { order = append(order, "last") }()
    }()

    want := []string{"last", "middle", "first"} // LIFO
    if !reflect.DeepEqual(order, want) {
        t.Fatalf("defer order = %v, want %v", order, want)
    }
}