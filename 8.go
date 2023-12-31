package main

import "fmt"

/*
Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
*/

// для примера будем разбирать с такими значениями(n = 3, pos = 2, bit = true (true - 1, false - 0))
func task8(n int64, pos uint, bit bool) (int64, error) {
	if pos > 63 {
		return 0, fmt.Errorf("invalid pos")
	}

	// конвертируем bool в 1 бит
	var bitValue uint8
	if bit {
		bitValue = 1
	}

	// делаем маску что бы удалить бит, который мы собираеся заменить
	var mask int64 = ^(int64(1) << pos)
	// Это операция нужна, для того, что если на позиции бит был 1 он стал 0
	// у нас выйдет такая таблица
	/*
	  +---+---+---+---+
	  | 0 | 0 | 1 | 1 | <- n
	  +---+---+---+---+
	  | 1 | 0 | 1 | 1 | <- mask
	  +---+---+---+---+
	  | 0 | 0 | 1 | 1 |
	  +---+---+---+---+
	*/
	n &= mask

	// и наконец добaвляем нашу потенциальную единичку
	/*
	  +---+---+---+---+
	  | 0 | 0 | 1 | 1 | <- n
	  +---+---+---+---+
	  | 0 | 1 | 0 | 0 | <- то что мы добaвляем
	  +---+---+---+---+
	  | 0 | 1 | 1 | 1 | -> 0 1 1 1 это 7, чего мы и ждали
	  +---+---+---+---+
	*/
	n |= int64(uint(bitValue) << pos)
	return n, nil
}
