package main

/*
К каким негативным последствиям может привести данный фрагмент кода, и как
это исправить? Приведите корректный пример реализации.

```go
var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}
func main() {
  someFunc()
}
```
*/

var justString string

func someFunc() {
	v := createHugeString(1 << 1)

	// если не добавить проверку, если создать слишком маленький string, мы получим панику, потому что слайс больше чем массив
	if len(v) > 100 {
		justString = v[:100]
	} else {
		justString = v
	}
}

func createHugeString(i int) string {
	return string(make([]byte, i))
}
