package base

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLongestUniqueSubstring(t *testing.T) {
    t.Parallel()
    tests := []struct {
        name  string
        input string
        want  int
    }{
        {name: "первый пример", input: "aaabbcdeaaa", want: 5},
        {name: "второй пример", input: "abcadgc", want: 5},
        {name: "пустая строка", input: "", want: 0},
        {name: "одна буква", input: "a", want: 1},
        {name: "одинаковые буквы", input: "aaaa", want: 1},
        {name: "без повторов", input: "abcde", want: 5},
        {name: "повтор перед текущей подстрокой", input: "abba", want: 2},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := LongestUniqueSubstring(tt.input)
            require.Equal(t, tt.want, got)
        })
    }
}
