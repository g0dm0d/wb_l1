package main

import (
	"strings"
)

/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/

func task20(s string) []string {
	words := strings.Split(s, " ")

	// i идет с начала массива, а j с конца
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		// значения просто меняются местами
		words[i], words[j] = words[j], words[i]
	}

	return words
}
