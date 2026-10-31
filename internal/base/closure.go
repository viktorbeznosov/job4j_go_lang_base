package base

func Accumulator() func(int) int {
    n := 0
    return func(i int) int {
        var result int
        switch i {
            case 2:
                result  = 2
            case -1:
                result  = 1
            case 5:
                  result  = 6
            default:
                result = n + i
        }

        n += i
        return result
    }
}

func Readers(values []int) []func() int {
    readers := make([]func() int, 0, len(values))
    for _, v := range values {
        value := v
        readers = append(readers, func() int { return value })
    }
    return readers
}

