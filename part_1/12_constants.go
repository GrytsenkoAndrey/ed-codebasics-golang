// Go: Константы

// Константы — это постоянные значения, которые инициализируются один раз и не изменяются в течение всего выполнения программы. В Go константы объявляются через ключевое слово const:

// const [название] [тип данных] = [значение]

// const StatusOk int = 200

// На практике тип данных не указывается, и несколько констант принято объявлять в рамках одного блока const:

// const (
//     StatusOk = 200
//     StatusNotFound = 404
// )

// Только некоторые типы данных можно присвоить константе: строки, символы, числа, логический тип:

// package main

// type Person struct {
// }

// func main() {
//     // такие константы допустимы
//     const (
//         num = 20
//         str = "hey"
//         isValid = true
//     )

//     // нельзя объявить структуру как константу
//     const p = Person{} // ошибка компиляции: const initializer Person{} is not a constant
// }

// Регистр первой буквы указывает на публичность/приватность константы:

// const (
//     // публичная константа, которую можно использовать во внешних пакетах
//     StatusOk = 200

//     // приватная константа, доступная только в рамках текущего пакета
//     statusInternalError = 500
// )

// Константы можно объявлять на уровне функции, либо пакета:

// package main

// import "fmt"

// const defaultStatus = 200

// func main() {
//     const status = 404

//     fmt.Println("default status:", defaultStatus) // default status: 200
//     fmt.Println("current status:", status) // current status: 404
// }

// Для последовательных числовых констант следует использовать идентификатор iota, который присвоит для списка чисел значения от его текущей позиции:

// package main

// import "fmt"

// const (
//     zero = iota
//     one
//     two
//     three
// )

// const (
// a = iota
// b = 42
// c = iota
// d
// )

// func main() {
// fmt.Println(zero, one, two, three) // 0 1 2 3
// fmt.Println(a, b, c, d)            // 0 42 2 3
// }

// Задание

// В больших проектах часто используется gRPC — высокопроизводительный RPC (Remote Procedure Call)-фреймворк для коммуникации между микросервисами. Ошибки в gRPC стандартизированы и представлены в виде строк для пользователя, либо в виде чисел для компьютера.

// Представим, что нам нужно написать конвертер ошибок в числовой формат для gRPC. Реализуйте функцию ErrorMessageToCode(msg string) int, которая возвращает числовой код для заданного значения. Список сообщений и соответствующих кодов:

// OK = 0
// CANCELLED = 1
// UNKNOWN = 2

// В реальности список намного шире. Мы для простоты ограничимся тремя ошибками. Учтите, что если в функцию передать неизвестную ошибку, она должна вернуть код ошибки для сообщения UNKNOWN.

package solution

const (
	OK = 0
	CANCELLED = 1
	UNKNOWN = 2
)

func ErrorMessageToCode(msg string) int {
	if msg == "OK" {
		return OK
	}
	if msg == "CANCELLED" {
		return CANCELLED
	}
	if msg != "OK" && msg != "CANCELLED" {
		return UNKNOWN
	}

	return UNKNOWN 
}