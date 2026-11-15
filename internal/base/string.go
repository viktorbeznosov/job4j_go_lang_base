package base

import "unicode/utf8"

func IsPalindrome(s string) bool {
    if !utf8.ValidString(s) {
        return false
    }

    isPalindrome := true

    runes := []rune(s)

    for i := 0; i < len(runes) - 1; i++ {
        if runes[i] != runes[len(runes) - i - 1] {
            isPalindrome = false
            break
        }
    }

    return isPalindrome
}


