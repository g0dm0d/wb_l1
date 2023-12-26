package main

import (
	"fmt"
	"sync"
)

/*
Дана последовательность чисел: 2,4,6,8,10. Найти сумму их
квадратов(2^2+3^2+4^2….) с использованием конкурентных вычислений.
*/

// наследует идею task_2
func task3() {
	arr := [5]int{2, 4, 6, 8, 10}

	c := make(chan int, len(arr))

	for _, i := range arr {
		go func(i int) {
			c <- i * i
		}(i)
	}

	sum := 0
	// как и в прошлом задание читаем из канала n раз и складываем
	for i := 0; i < len(arr); i++ {
		sum += <-c
	}
	fmt.Println(sum)
}

// наследует идею task_2_2
func task3_2() {
	arr := [5]int{2, 4, 6, 8, 10}

	// заранее инициализируем для наши рутин переменную в которую будут все складывать
	sum := 0

	wg := new(sync.WaitGroup)

	for _, i := range arr {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			// в рутине считаем
			sum += i * i
		}(i)
	}

	wg.Wait()
	fmt.Println(sum)
}
