package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

/*
Реализовать постоянную запись данных в канал (главный поток). Реализовать
набор из N воркеров, которые читают произвольные данные из канала и
выводят в stdout. Необходима возможность выбора количества воркеров при старте.
*/

const WORKERS = 5

func task4() {
	mainChannel := make(chan string)

	wg := new(sync.WaitGroup)

	c := make(chan os.Signal)
	// Создается хендлер сигнала прерывания
	// который при срабатывание отправляет сигнал в канал
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Запускаем воркеры
	for i := 1; i <= WORKERS; i++ {
		wg.Add(1)
		go worker(mainChannel, wg)
	}

	// В отдельном потоке запускаю считывания stdin
	go func() {
		for {
			var data string
			fmt.Scanln(&data)
			// отправляю данные в общий канал
			mainChannel <- data
		}
	}()

	// мейн поток ждет когда придет сигнал прерывания
	<-c
	// закрываю канал, что бы отправить в потоки ничего нельзя было и они сами закрылись после того как выполнят задачу
	close(mainChannel)
	// жду когда потоки даделают работу
	wg.Wait()
}

// Воркер, который читает данные из канала и выводит их
func worker(ch <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		data, ok := <-ch
		// Если if отрабатывает значит данные в канале были все прочитаны и он закрыт
		if !ok {
			return
		}

		fmt.Println(data)
	}
}
