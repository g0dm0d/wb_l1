package main

import "sync"

/*
Реализовать конкурентную запись данных в map.
*/

func task7() {
	mymap := make(map[int]int)
	wg := new(sync.WaitGroup)
	// Создаем мьютекс
	mu := new(sync.Mutex)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			// блокируем мьютекс, что бы он был доступен только в это рутине
			mu.Lock()

			// деферим что бы по завершению функции блокировка снялась и мапу можно было изменять в другой рутине
			defer mu.Unlock()
			defer wg.Done()

			mymap[1]++
		}()
	}
	wg.Wait()
}
