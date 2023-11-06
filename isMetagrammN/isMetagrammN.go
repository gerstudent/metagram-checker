package main

// Функция для проверки того, является ли строка t метаграммой
// строки s, в которой заменили n букв.
func isMetagrammN(s, t string, n int) bool {

	// Преобразование строк в слайс рун для корректной обработки русских символов.
	sRune := []rune(s)
	tRune := []rune(t)

	// Проверка равенства длин строк и корректности значения n.
	if len(s) != len(t) || n > len(s) {
		return false
	}

	// Создание счетчика отличающихся значений.
	var differenceCount int

	for i := 0; i < len(sRune); i++ {

		// Если текущие коды символов не равны, то увеличиваем счетчик.
		if sRune[i] != tRune[i] {
			differenceCount++
		}
	}

	// Если отличаются ровно n символов, вернем true. Иначе - false.
	return differenceCount == n
}
