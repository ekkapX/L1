package main

/*
Рассмотреть следующий код и ответить на вопросы: к каким негативным последствиям он может привести и как это исправить?

Приведите корректный пример реализации.

var justString string

func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}

Вопрос: что происходит с переменной justString?

justString удерживает неиспользуемую память от сборщика мусора.

Вся строка остается в памяти, даже если мы взяли только первые 100 байт. Сборщик мусора не сможет почистить память,
т.к строка ссылается на первоначальную строку из 1024 байт. Можно сделать копию строки в новый слайс
*/

var justString string

func createHugeString(n int) string {
	return string(make([]byte, n))
}

func someFunc() {
	v := createHugeString(1 << 10)
	buf := make([]byte, 100)
	copy(buf, v[:100])
	justString = string(buf)
}

func main() {
	someFunc()
}
