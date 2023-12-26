package main

import "time"

/*
Реализовать собственную функцию sleep.
*/

func task25() {
	sleep(1000)
	println("hello")
}

func sleep(x int) {
	// ждем x миллисекунд что бы получить данные из канала, после чего функция выполнится и код продолжит выполнятся
	<-time.After(time.Millisecond * time.Duration(x))
}
