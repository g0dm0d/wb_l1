package main

/*
Разработать программу, которая переворачивает подаваемую на ход строку
(например: «главрыба — абырвалг»). Символы могут быть unicode.
*/

func task19(input string) string {
	runes := []rune(input)

	// i идет с начала массива, а j с конца
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		// значения просто меняются местами
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
