package main

import "sort"

/*
Реализовать бинарный поиск встроенными методами языка.
*/

func task17() {
	arr := []int{1, 2, 3, 4, 5}
	target := 2

	index := sort.Search(len(arr), func(i int) bool { return arr[i] >= target })
	println(index)
}
