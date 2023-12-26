package main

import "fmt"

/*
Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из
массива, во второй — результат операции x*2, после чего данные из второго
канала должны выводиться в stdout.
*/

func task9() {
	arr := [5]int{2, 4, 6, 8, 10}

	in := make(chan int)
	out := make(chan int)

	// проходимя по массиву и отправляем все в канал in
	go func() {
		for _, i := range arr {
			in <- i
		}
		// закрываем, что бы не было дедлока
		close(in)
	}()

	go func() {
		// читаем как из массива т.к. in закроется
		for v := range in {
			out <- v * 2
		}
		// закрываем, что бы не было дедлока
		close(out)
	}()

	// опять читаем
	for v := range out {
		fmt.Println(v)
	}
}