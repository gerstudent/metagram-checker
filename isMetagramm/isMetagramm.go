package main

// Функция для проверки того, является ли строка t метаграммой строки s.
func isMetagramm(s, t string) bool {

	// Преобразование строк в слайс рун для корректной обработки русских символов.
	sRune := []rune(s)
	tRune := []rune(t)

	// Проверка равенства длин строк.
	if len(sRune) != len(tRune) {
		return false
	}

	// Создание счетчика отличающихся значений.
	var differenceCount int

	for index := 0; index < len(sRune); index++ {

		// Если текущие коды символов не равны, то увеличиваем счетчик.
		if sRune[index] != tRune[index] {
			differenceCount++
		}
	}

	// Если отличается только один символ, вернем true. Иначе - false.
	return differenceCount == 1
}
