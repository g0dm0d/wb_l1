package main

import "fmt"

func task26() {
	fmt.Println(IsUnique("abcd"))
}

func IsUnique(s string) bool {
	charMap := make(map[rune]bool)

	// прохожусь по всем символам
	for _, char := range []rune(s) {
		// если в хешмапе уже есть эта руна, значит это повторение и строка не все символы уникальные
		if _, ok := charMap[char]; ok {
			return false
		}
		charMap[char] = true
	}
	return true
}
