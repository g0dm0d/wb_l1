## 1. Какой самый эффективный способ конкатенации строк?

Builder из пакета strings

## 2. Что такое интерфейсы, как они применяются в Go?

Интерфейс это тип для абстракции данных.

## 3. Чем отличаются RWMutex от Mutex?

RWMutex - позвляет нескольким горутинам читать, а Mutex позволяет только 1 горутине.
RWmutex есть деление на писателя и читателя. В Mutex его нет

## 4. Чем отличаются буферизированные и не буферизированные каналы?

В не буферизированном канале делается пауза, пока из канал не прочтут данные.
В буферизированном канале делается пауза, только если он заполнен

## 5. Какой размер у структуры struct{}{}

0

## 6. Есть ли в Go перегрузка методов или операторов?

нет

## 7. В какой последовательности будут выведены элементы map[int]int?

зависит от того как выводить. Если через fmt.Println то он отсортирует, если например через range - тогда в случайном

## 8. В чем разница make и new?

make - инициализирует, new - выделяем место в памяти. new используют для структур, а make для каналов, hashmap

## 9. Сколько существует способов задать переменную типа slice или map?

- через make

```go
m := make(map[string]int)
```

- через объявление переменной

```
var m1 []int

m2 := map[string]int{
    "one":   1,
}
```

## 10. Что выведет данная программа и почему?

1 1 - потому что это копия указателя которая живет только внутри функции update и на main не влияет


```go
func update(p *int) {
	b := 2
	p = &b
}
func main() {
	var (
		a = 1
		p = &a
	)
	fmt.Println(*p)
	update(p)
	fmt.Println(*p)
}

```

## 11. Что выведет данная программа и почему?

```go
func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg sync.WaitGroup, i int) {
			fmt.Println(i)
			wg.Done()
		}(wg, i)
	}
	wg.Wait()
	fmt.Println("exit")
}
```

дедлок потому что в wg передается без указателя и никогда не обнулится


## 12 Что выведет данная программа и почему?

```go
func main() {
	n := 0
	if true {
		n := 1
		n++
	}
	fmt.Println(n)
}
```

0 - потому что n которая внутри if это другая переменная

## 13 Что выведет данная программа и почему?

```go
func someAction(v []int8, b int8) {
	v[0] = 100
	v = append(v, b)
}

func main() {
	var a = []int8{1, 2, 3, 4, 5}
	someAction(a, 6)
	fmt.Println(a)
}
```

100, 2, 3, 4, 5

потому что после append это уже другой slice, а 1ое значение изменили до создания нового слайса

## 14 Что выведет данная программа и почему?

```go
func main() {
	slice := []string{"a", "a"}
	func(slice []string) {
		slice = append(slice, "a")
		slice[0] = "b"
		slice[1] = "b"
		fmt.Print(slice)
	}(slice)
	fmt.Print(slice)
}
```

bba, aa

потому что 1 нас 2 разных слайса, есть слайс внутри функции который будет равен bba и в main который не изменился, потому что первая операции внутри функции создает новый слайс
