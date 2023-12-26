package main

/*
Удалить i-ый элемент из слайса.
*/

func task23[T any](arr *[]T, ind int) {
	// Перезапись слайса без i-ыйного элемента
	*arr = append((*arr)[:ind], (*arr)[ind+1:]...)
}
