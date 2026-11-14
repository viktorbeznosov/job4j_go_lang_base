package base

import "strings"

type Set map[rune]struct{}

func (s Set) Add(value rune) { s[value] = struct{}{} }

func (s Set) AddString(value string) {
    for _, r := range value {
        s.Add(r)
    }
}

func (s Set) Contains(value rune) bool {
    _, ok := s[value]
    return ok
}

func LongestUniqueSubstring(value string) int {
    strArr := strings.Split(value, " ")

    maxLength := 0
    for _, item := range strArr {
        strSet := Set{}
        strSet.AddString(item)
        if length := len(strSet);length > maxLength {
            maxLength = length
        }
    }

    return maxLength
}



