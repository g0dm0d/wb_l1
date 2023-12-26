package main

import (
	"fmt"
	"math/big"
)

/*
Разработать программу, которая перемножает, делит, складывает, вычитает две
числовых переменных a,b, значение которых > 2^20.
*/

// Все делаю через стандартную библиотеку golang
func task22() {
	a := new(big.Int).Lsh(big.NewInt(1), 21)
	b := new(big.Int).Lsh(big.NewInt(1), 21)

	var mul, div, add, sub big.Int
	mul.Mul(a, b)
	div.Div(a, b)
	add.Add(a, b)
	sub.Sub(a, b)

	fmt.Println(a.String(), b.String(), mul.String(), div.String(), add.String(), sub.String())
}
