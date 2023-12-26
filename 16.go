package main

import (
	"fmt"
	"sort"
)

/*
Реализовать быструю сортировку массива (quicksort) встроенными методами
языка.
*/

func task16() {
	arr := []int{5, 4, 3, 2, 1}

	sort.Ints(arr)

	fmt.Println(arr)
}
