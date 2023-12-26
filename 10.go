package main

/*
Дана последовательность температурных колебаний:  -27.0 13.0, 19.0,
15.5, 24.5, -21.0, 32.5. Объединить данные значения в группы с шагом в 10
градусов. Последовательность в подмножноствах не важна.
*/

func task10() {
	degree := [7]float32{-27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	degreeSorted := make(map[int][]float32)

	for _, val := range degree {
		// делем, что бы убрать 1ый десяток, затем умнажаем на 0 что бы получить ровное число
		// было 15.5 -> 1 * 10 -> 10
		step := int(val/10) * 10
		// и просто сохраняем
		degreeSorted[step] = append(degreeSorted[step], val)
	}
}
