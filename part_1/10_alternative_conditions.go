// x := 10

// switch x {
//     default: // default всегда выполняется последним независимо от расположения в конструкции
//         fmt.Println("default case")
//     case 10:
//         fmt.Println("case 10")
// }

// Однако при необходимости можно реализовать логику С-подобных языков и «провалиться» в следующий case:

// x := 10

// switch { // выражение отсутствует. Для компилятора выглядит как: switch true
// 	default:
// 		fmt.Println("default case")
// 	case x == 10:
// 		fmt.Println("equal 10 case")
// 		fallthrough
// 	case x <= 10:
// 		fmt.Println("less or equal 10 case")
// }


// Для замены символов в строке существует функция ReplaceAll(s, old, new string) string из пакета strings:
// strings.ReplaceAll("hello world!", "world!", "buddy!") // hello buddy!
// strings.ReplaceAll("one two three", " ", "_") // one_two_three


// Go: Альтернативная условная конструкция

// В Go присутствует единственная альтернатива if — конструкция switch. Для этой конструкции используется стандартный синтаксис, но логика работы отличается от С-подобных языков. Когда срабатывает условие какого-либо case, программа выполняет блок и выходит из конструкции switch без необходимости писать break:

// x := 10

// switch x {
//     default: // default всегда выполняется последним независимо от расположения в конструкции
//         fmt.Println("default case")
//     case 10:
//         fmt.Println("case 10")
// }

// Output:

// case 10

// Однако при необходимости можно реализовать логику С-подобных языков и «провалиться» в следующий case:

// x := 10

// switch { // выражение отсутствует. Для компилятора выглядит как: switch true
//     default:
//         fmt.Println("default case")
//     case x == 10:
//         fmt.Println("equal 10 case")
//         fallthrough
//     case x <= 10:
//         fmt.Println("less or equal 10 case")
// }

// Output:

// equal 10 case
// less or equal 10 case

// Задание

// Реализуйте функцию ModifySpaces(s, mode string) string, которая изменяет строку s в зависимости от переданного режима mode.


package solution

import (
	"strings"
)

func ModifySpaces(s, mode string) string {
	switch {
		default:
			return strings.ReplaceAll(s, " ", "*")
		case mode == "dash":
			return strings.ReplaceAll(s, " ", "-")
		case mode == "underscore":
			return strings.ReplaceAll(s, " ", "_")
		case mode == "underscore":
			return strings.ReplaceAll(s, " ", "*")
	}
}