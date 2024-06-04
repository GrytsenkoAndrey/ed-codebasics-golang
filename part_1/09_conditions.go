// Go: Условные конструкции

// Условия в Go представлены привычной конструкцией if else. В условии должно быть строго выражение логического типа. Следующий пример вернет ошибку компиляции:

// if "hi" { // non-bool "hi" (type string) used as if condition
// }

// Корректный пример:

// package main

// import (
//     "fmt"
//     "strings"
// )

// func statusByName(name string) string {
//     // функция проверяет, что строка name начинается с подстроки "Mr."
//     if strings.HasPrefix(name, "Mr.") {
//         return "married man"
//     } else if strings.HasPrefix(name, "Mrs.") {
//         return "married woman"
//     } else {
//         return "single person"
//     }
// }

// func main() {
//     n := "Mr. Doe"
//     fmt.Println(n + " is a " + statusByName(n)) // Mr. Doe is a married man

//     n = "Mrs. Berry"
//     fmt.Println(n + " is a " + statusByName(n)) // Mrs. Berry is a married woman

//     n = "Karl"
//     fmt.Println(n + " is a " + statusByName(n)) // Karl is a single person
// }

// Логическое выражение пишется после if без скобок. else if можно написать только раздельно.
// Задание

// На веб-сайтах часто используются разные поддомены для языков. Например, сайт site.com на английском располагается по адресу en.site.com, а на русском — ru.site.com.

// Реализуйте функцию DomainForLocale(domain, locale string) string, которая добавляет язык locale как поддомен к домену domain. Язык может прийти пустым, тогда нужно добавить поддомен en.. 




// package main

// import (
//     "fmt"
//     "strings"
// )

// func statusByName(name string) string {
//     // функция проверяет, что строка name начинается с подстроки "Mr."
//     if strings.HasPrefix(name, "Mr.") {
//         return "married man"
//     } else if strings.HasPrefix(name, "Mrs.") {
//         return "married woman"
//     } else {
//         return "single person"
//     }
// }

// func main() {
//     n := "Mr. Doe"
//     fmt.Println(n + " is a " + statusByName(n)) // Mr. Doe is a married man

//     n = "Mrs. Berry"
//     fmt.Println(n + " is a " + statusByName(n)) // Mrs. Berry is a married woman

//     n = "Karl"
//     fmt.Println(n + " is a " + statusByName(n)) // Karl is a single person
// }

package solution

import (
	"fmt"
)

func DomainForLocale(domain, locale string) string {
	result := ""
	if len(locale) == 0 {
		result = fmt.Sprintf("en.%s", domain)
	} else {
		result = fmt.Sprintf("%s.%s", locale, domain)
	}

	return result
}

