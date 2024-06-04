package solution

func IsValid(id int, text string) bool {
	return id > 0 && text != ""
}

// Go: Логический тип

// Логический тип в Go представлен привычными значениями true и false c операторами:
// — && (и)
// — == (равно)
// — || (или)
// — ! (не):

// true && false // false
// false || true // true

// Объявление переменных происходит через ключевое слово bool:

// var b bool = true

// // сокращенная запись
// bs := false

// Из-за строгой типизации в Go можно сравнивать только одинаковые типы данных:

// true == false // false

// false == false // true

// То есть нельзя сравнить строку с логическим типом, как это происходит в динамических языках:

// true == "hello" // invalid operation: false == "hello" (mismatched types untyped bool and untyped string)

// Когда необходимо проверить на истинность разные значения, нелогические типы нужно привести к логическому:

// flag := true
// text := "hello"

// // вариант не сработает, потому что нельзя конвертировать строку в bool
// flag && bool(text) // cannot convert text (type string) to type bool

// // правильный вариант: если строка не пустая, то в ней есть текст
// flag && text != "" // true

// Задание

// Реализуйте функцию IsValid(id int, text string) bool, которая проверяет, что переданный идентификатор id является положительным числом и текст text не пустой.
// Например:

// IsValid(0, "hello world") // false
// IsValid(-22, "hello world") // false
// IsValid(22, "") // false
// IsValid(225, "hello world") // true
