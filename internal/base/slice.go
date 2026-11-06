package base

func RemoveValue(items []int, value int) []int {
    write := 0
    for read := 0; read < len(items); read++ {
        if items[read] == value {
            continue
        }
        items[write] = items[read]
        write++
    }
    clear(items[write:])
    return items[:write]
}

func KeepPositive(items []int) []int {
        write := 0
        for read := 0; read < len(items); read++ {
            if items[read] <= 0 {
                continue
            }
            items[write] = items[read]
            write++
        }
        clear(items[write:])
        return items[:write]
}



