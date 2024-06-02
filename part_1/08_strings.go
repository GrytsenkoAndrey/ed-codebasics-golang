//// обрезает символы, переданные вторым аргументом, с обеих сторон строки
//Trim(s, cutset string) string
//// пример
//strings.Trim(" hello ", " ") // "hello"

//// преобразует все буквы в строке в нижний регистр
//strings.ToLower(s string) string
//// пример
//strings.ToLower("пРиВеТ") // "привет"

//// озаглавливает первую букву в каждом слове в строке
//strings.Title(s string) string
//// пример
//strings.Title("привет, джон") // "Привет, Джон"

package solution

import (
	"fmt"
	"strings"
)

func Greetings(name string) string {
	return fmt.Sprintf("Привет, %s!", strings.Title(strings.ToLower(strings.Trim(name, " "))))
}