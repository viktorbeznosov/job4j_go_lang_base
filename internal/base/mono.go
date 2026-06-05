package base

func Mono(nums []int) bool {
	if len(nums) <= 2 {
		return true
	}

	length := len(nums)
	for i := 1; i < length-1; i++ {
		// Проверяем, является ли элемент локальным минимумом или максимумом
		prev, curr, next := nums[i-1], nums[i], nums[i+1]

		// Локальный максимум: prev < curr > next
		// Локальный минимум: prev > curr < next
		if (prev < curr && curr > next) || (prev > curr && curr < next) {
			return false
		}
	}

	return true
}
