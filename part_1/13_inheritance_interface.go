// Go: Наследование и интерфейсы

//     Композиция
//     Встраивание
//     Интерфейсы

// В языке Go отсутствует понятие наследования в классическом виде (нет ключевого слова extends, как, например, в Java).
// Однако терять преимущества, которые даёт механизм наследования, разработчики Go не хотели.
// Поэтому всё тоже самое, что можно сделать в других языках программирования за счёт наследования классов, можно реализовать в Golang другими средствами.

// Для начала определим, что даёт нам наследование:
// 1. Повторное использование кода: класс-наследник получает всё содержимое класса-предка и добавляет своё.
// 2. Динамический полиморфизм: переменная, у которой типом данных является некий базовый класс, может ссылаться как на объекты этого базового класса, так и на объекты его класса-наследника.
// 3. Динамическая диспетчеризация: Метод с одним и тем же названием может иметь разную реализацию в классе-предке и классе-наследнике.

// В качестве элегантного решения проблемы повторного использования кода Golang предлагает использование композиции и встраивания.
// Функционал полиморфизма и динамической диспетчеризации достигается за счёт использования интерфейсов.
// Композиция

// Рассмотрим пример использования композиции в качестве инструмента повторного использования кода.
// Допустим мы имеем структуру, которая описывает Машину (Сar).
// Если нам необходимо получить все возможности структуры Car и дополнить их в классе Пожарная машина (FireEngine), то мы можем использовать композицию (сделать FireEngine членом Car):

// type Car struct {
//   // … содержимое
// }

// type FireEngine struct {
//     basis Car
//     // … дополнение
// }

// Встраивание

// Теперь рассмотрим решение проблемы повторного использования кода через Встраивание.
// Допустим структура Car имеет метод Drive. Мы должны скопировать точное поведение метода Drive в структуре FireEngine.
// Для этого мы можем применить делегирование:

// type Car struct {
//   // … содержимое
// }

// func (c *Car) Drive() { … }

// type FireEngine struct {
//     basis Car
//     // … дополнение
// }

// func (fe *FireEngine) Drive() { fe.basis.Drive() }

// Однако оно ведёт к дублированию кода. Поэтому имеет механизм Встраивание, что позволяет значительно сократить код:

// type Car struct {
//   // … содержимое
// }

// func (c *Car) Drive() { … }

// type FireEngine struct {
//     Car
//     // … дополнение
// }

// Интерфейсы

// Теперь на очереди функционал полиморфизма и динамической диспетчеризации.
// Допустим, что наше приложение расширяется и в ней появляется всё больше видов специализированных машин:
// Полицейская Машина (PoliceCar), Машина Скорой Помощи (AmbulanceCar), Поливомоечная машина (WateringCar).
// Все они должны иметь метод Drive, однако реализует его каждая по-разному.
// Например, PoliceCar едет со звуком сирены, а WateringCar во время поездки поливает дорогу водой.
// То есть, мы должны определить "поведение", которое должно присутствовать в каждой из этих структур, но реализовано оно может быть по-разному.
// В таком случае на сцену и выходят интерфейсы (interfaces).

// Интерфейсы определяют, что тип делает, а не кем он является.
// Методы должны отражать поведение типа, поэтому интерфейсы объявляются с набором методов, которые тип должен обязательно иметь (-able).
// В нашем случае каждая из указанных выше структур должна иметь метод Drive.

// type IDriveable interface {
//   Drive()
// }

// type Car struct {
//   // … 
// }

// type PoliceCar struct {
//     // … 
// }

// func (c Car) Drive() {
//     fmt.Println("Просто еду по дороге")
// }

// func (pc PoliceCar) Drive() {
//     fmt.Println("Еду по дороге с мигалкой. Виу-виу!")
// }

// func main() {
//   cars := []IDriveable{&Car{}, &PoliceCar{}}
//   for _, vehicle := range cars {
//       vehicle.Drive()
//       // => Просто еду по дороге
//       // =>  Еду по дороге с мигалкой. Виу-виу!
//   }
// }

// Именование интерфейсов в виде "глагол + able" стандартно для большинства языков программирования.
// Однако в Go интерфейсы именуются немного по-другому. В данном случае интерфейс должен называться Driver.
// Подробнее про нейминг можно почитать в официальной документации Golang.

// Так никакого явного указание реализации не требуется.
// Любой тип, который предоставляет методы, которые указаны в интерфейсе, можно считать реализующим интерфейс.
// Задание

// Реализуйте интерфейс IVoiceable для структур Cat, Cow и Dog так, чтобы при вызове метода Voice экземпляр структуры Cat возвращал строку "Мяу", экземпляр Cow строку "Мууу", а экземпляр Dog сообщение Гав:

// cat := Cat{} 
// dog := Dog{}
// cow := Cow{}

// fmt.Println(cat.Voice()) // Мяу
// fmt.Println(dog.Voice()) // Гав
// fmt.Println(cow.Voice()) // Мууу

package solution

type Voicer interface {
    Voice() string
}

type Cat struct {
    // … 
}

type Cow struct {
    // … 
}

type Dog struct {
	// … 
}

