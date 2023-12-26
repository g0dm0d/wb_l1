package main

import (
	"fmt"
	"sync"
)

/*
Реализовать структуру-счетчик, которая будет инкрементироваться в
конкурентной среде. По завершению программа должна выводить итоговое значение счетчика.
*/

type Atomic struct {
	count int

	mu sync.Mutex
}

func NewAtomic() *Atomic {
	return &Atomic{
		count: 0,
	}
}

func (a *Atomic) CounterInc() {
	// блокирю мьютекс для другиз рутин
	a.mu.Lock()

	// деферю разблокировку
	defer a.mu.Unlock()

	// инкременчу на 1
	a.count++
}

func task18() {
	atomic := NewAtomic()

	wg := new(sync.WaitGroup)

	// запускаем рутины
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 100; i++ {
				atomic.CounterInc()
				wg.Done()
			}
		}()
	}

	wg.Wait()
	fmt.Println(atomic.count)
}
