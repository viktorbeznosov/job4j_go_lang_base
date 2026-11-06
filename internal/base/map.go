package base

import "slices"

func Frequencies(words []string) map[string]int {
    frequencies := make(map[string]int)

    for _, value := range words {
        count := frequencies[value]
        count++
        frequencies[value] = count
    }

    return frequencies
}

func MostFrequent(freq map[string]int) (word string, count int, ok bool) {
    values := make([]int, 0, len(freq))
    for _, value := range freq {
        values = append(values, value)
    }

    if len(values) == 0 {
        return "", 0, false
    }

    max := slices.Max(values)
    result := ""

    for key := range freq {
        if freq[key] == max && (len(result) == 0 || len(result) > 0 && len(result) > len(key)) {
            result = key
        }
    }

    return result, max, true
}


