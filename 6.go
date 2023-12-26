package main

import (
	"context"
	"runtime"
)

/*
Реализовать все возможные способы остановки выполнения горутины.
*/

// при остановки main, все рутины умeрают
func task6() {
	go func() {}()
}

// использоание функции из пакета runtime которая убивает все рутины
func task6_2() {
	go func() {
		runtime.Goexit()
	}()
}

// сигнал через канал
func task6_3() {
	c := make(chan bool, 1)

	go func() {
		select {
		case <-c:
			return
		}
	}()

	c <- true
}

// через контекст
func task6_4() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-ctx.Done()

	}()
	cancel()
}
