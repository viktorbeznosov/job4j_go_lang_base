package base

import (
	"fmt"
	"strconv"
)

func ParseCelsius(raw string) (int16, error) {
	n, err := strconv.ParseInt(raw, 10, 16)
	if err != nil {
		return 0, fmt.Errorf("parsing error: %w", err)
	}

	const (
		minInt16 = -1 << 15
		maxInt16 = 1<<15 - 1
	)
	if n < minInt16 || n > maxInt16 {
		return 0, fmt.Errorf("value %d is outside int16 range", n)
	}

	return int16(n), nil
}

func ToFahrenheit(c int16) float64 {
	// int16 overflows on c * 9 at the type bounds; widen first.
	result := int64(c)*9/5 + 32
	return float64(result)
}
