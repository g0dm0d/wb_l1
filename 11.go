package main

import "fmt"

/*
Реализовать пересечение двух неупорядоченных множеств.
*/

func task11() {
	arr1 := [5]int{2, 4, 6, 8, 10}
	arr2 := [5]int{3, 5, 6, 9, 10}

	// создаю хешмапу что бы в ней хранить числа из первого массива
	setMap := make(map[int]bool)

	for _, v := range arr1 {
		setMap[v] = true
	}

	res := []int{}

	for _, v := range arr2 {
		// если ok то значит значения есть в обоих массивах
		if _, ok := setMap[v]; ok {
			res = append(res, v)
		}
	}

	fmt.Println(res)
}
