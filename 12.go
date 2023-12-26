package main

import "fmt"

/*
Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее
собственное множество.
*/

func task12() {
	strings := []string{"cat", "cat", "dog", "cat", "tree"}
	setmap := make(map[string]bool)

	// из-за того, что ключи в мапе не повторяются, у нас получается ножество
	for _, v := range strings {
		setmap[v] = true
	}

	for k := range setmap {
		fmt.Println(k)
	}
}
