package base

import "testing"

func TestIsPalindrome(t *testing.T) {
    t.Parallel()
    tests := []struct {
        name string
        s    string
        want bool
    }{
        {"level", "level", true},
        {"levvel", "levvel", true},
        {"Go", "Go", false},
        {"spaces", "а роза упала на лапу азора", false},
        {"топот", "топот", true},
        {"case", "Аа", false},
        {"empty", "", true},
        {"invalid UTF-8", "\xff", false},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            t.Parallel()
            got := IsPalindrome(tt.s)
            if got != tt.want {
                t.Errorf("IsPalindrome(%q) = %v, want %v", tt.s, got, tt.want)
            }
        })
    }
}