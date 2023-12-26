package main

import (
	"fmt"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения в
канал, а с другой стороны канала — читать. По истечению N секунд программа
должна завершаться.
*/

func SecondChann(c chan<- int, n int) {
	// Пользуюсь встроенной функцией в пакете time, которая отправит сигнал в канал, когда n секунд пройдет
	timeout := time.After(time.Duration(n) * time.Second)
	for {
		select {
		// жду сигнал, что бы закрыть канал и программа закончила работу
		case <-timeout:
			close(c)
			return
		default:
			c <- n
		}
	}
}

func task5() {
	c := make(chan int)

	go SecondChann(c, 5)

	for v := range c {
		fmt.Println(v)
	}
}
