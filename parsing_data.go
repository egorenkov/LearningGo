package main

import (
	"fmt"
	"regexp"
)

func main() {

	str := "name 30, 24,,11"
	re := regexp.MustCompile(`\name+`)

	s := re.FindString(str)
	fmt.Print(s)

}

/*
Из строки в число
i, err := strconv.Atoi("123")
f, err := strconv.ParseFloat("3.14", 64)

Из Число -> строка
s := strconv.Itoa(123)
s := strconv.FormatFloat(3.14, 'f', 2, 64)


Мы задаем регулярное выражение
	re := regexp.MustCompile(`\d+`)

Мы проверяем подходит ли строка под данное выражение
	s := re.MatchString(str)

Для нахождения первого совпадения
	s := re.FindString(str)

Для нахождения групп всех совпадния
	s := re.FindAllString(str, -1)

Для нахождения групп определенного количества совпадния
	s := re.FindAllString(str, 2)

Для обычной замены
	s := re.ReplaceAllString(str, "1")

Захват групп

re := regexp.MustCompile(`(\d+)-(\d+)-(\d+)`)
matches := re.FindStringSubmatch("2023-11-15")

fmt.Println(matches[0]) // "2023-11-15" (полное совпадение)
fmt.Println(matches[1]) // "2023" (первая группа)
fmt.Println(matches[2]) // "11" (вторая группа)
fmt.Println(matches[3]) // "15" (третья группа)

Поиск чисел
numRegex := regexp.MustCompile(`-?\d+\.?\d*`)

Поиск дат
dateRegex := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`) // YYYY-MM-DD


Используйте raw-строки (`) для regexp чтобы избежать двойного экранирования
*/

/*
i, err := strconv.Atoi("123")
f, err := strconv.ParseFloat("3.14", 64)

// Число -> строка
s := strconv.Itoa(123)
s := strconv.FormatFloat(3.14, 'f', 2, 64)


libary strings
Поиск подстроки
	strings.Contains("hello world", "wor")  // true

Разделение по разделителю
	strings.Split("a,b,c", ",")  // ["a", "b", "c"]

Соединяет строки
	strings.Join([]string{"a", "b"}, "-")  // "a-b"

В нижний регистр
	strings.ToLower("HELLO")  // "hello"

В верхний регистр
	strings.ToUpper("hello")  // "HELLO"

Обрезает указанные символы
	strings.Trim("!!!hello!!!", "!")  // "hello"

Заменяте подстроки
	Replace(s, old, new string, n int) string - заменяет подстроки

Заменяет все вхождения
	ReplaceAll(s, old, new string) string


*/

/*
Работа с регулярными выражениями в Go: полное руководство
Go предоставляет мощный пакет regexp для работы с регулярными выражениями. Рассмотрим все полезные методы и примеры.

Проверка соответствия строки регулярному выражению
1. Простая проверка
go
package main

import (
    "fmt"
    "regexp"
)

func main() {
    matched, err := regexp.MatchString(`foo.*`, "seafood")
    fmt.Println(matched, err) // true <nil>

    matched, err = regexp.MatchString(`bar.*`, "seafood")
    fmt.Println(matched, err) // false <nil>
}
2. Скомпилированное выражение (для многократного использования)
go
re := regexp.MustCompile(`foo.*`)
fmt.Println(re.MatchString("seafood")) // true
fmt.Println(re.MatchString("barmenu")) // false
Извлечение частей строки
1. Поиск первого совпадения
go
re := regexp.MustCompile(`\d+`)
fmt.Println(re.FindString("123 abc 456")) // "123"
2. Поиск всех совпадений
go
re := regexp.MustCompile(`\d+`)
fmt.Println(re.FindAllString("123 abc 456 def 789", -1)) // ["123" "456" "789"]
3. Ограничение количества совпадений
go
re := regexp.MustCompile(`\d+`)
fmt.Println(re.FindAllString("123 abc 456 def 789", 2)) // ["123" "456"]
Работа с группами захвата
1. Извлечение групп
go
re := regexp.MustCompile(`(\d+)-(\d+)-(\d+)`)
matches := re.FindStringSubmatch("2023-11-15")

fmt.Println(matches[0]) // "2023-11-15" (полное совпадение)
fmt.Println(matches[1]) // "2023" (первая группа)
fmt.Println(matches[2]) // "11" (вторая группа)
fmt.Println(matches[3]) // "15" (третья группа)
2. Именованные группы
go
re := regexp.MustCompile(`(?P<year>\d+)-(?P<month>\d+)-(?P<day>\d+)`)
matches := re.FindStringSubmatch("2023-11-15")

for i, name := range re.SubexpNames() {
    if i != 0 && name != "" {
        fmt.Printf("%s: %s\n", name, matches[i])
    }
}
// Вывод:
// year: 2023
// month: 11
// day: 15
Замена в строках
1. Простая замена
go
re := regexp.MustCompile(`a(x*)b`)
fmt.Println(re.ReplaceAllString("-ab-axxb-", "T")) // "-T-T-"
2. Замена с использованием групп
go
re := regexp.MustCompile(`(\w+) (\w+)`)
fmt.Println(re.ReplaceAllString("John Doe", "$2, $1")) // "Doe, John"
3. Замена с функцией-обработчиком
go
re := regexp.MustCompile(`[^aeiou]`)
fmt.Println(re.ReplaceAllStringFunc("hello world", strings.ToUpper))
// "heLLo woLLD"
Разделение строк
go
re := regexp.MustCompile(`a(x*)b`)
fmt.Printf("%q\n", re.Split("-ab-axxb-", -1)) // ["-" "-" "-"]
Полезные шаблоны регулярных выражений
1. Проверка email
go
emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
fmt.Println(emailRegex.MatchString("test@example.com")) // true
2. Поиск URL
go
urlRegex := regexp.MustCompile(`https?://[^\s]+`)
fmt.Println(urlRegex.FindAllString("Visit https://example.com", -1)) // ["https://example.com"]
3. Извлечение хэштегов
go
tagRegex := regexp.MustCompile(`#\w+`)
fmt.Println(tagRegex.FindAllString("#golang is #awesome", -1)) // ["#golang" "#awesome"]
4. Поиск чисел
go
numRegex := regexp.MustCompile(`-?\d+\.?\d*`)
fmt.Println(numRegex.FindAllString("123 -3.14 0.5", -1)) // ["123" "-3.14" "0.5"]
Оптимизация производительности
Компилируйте регулярные выражения один раз:

go
var re = regexp.MustCompile(`pattern`) // Глобальная переменная
Используйте FindAllStringSubmatch вместо нескольких вызовов:

go
// Неэффективно:
re.FindString(...)
re.FindString(...)

// Эффективно:
re.FindAllString(..., -1)
Для простых операций используйте strings:

go
// Вместо regexp для простого поиска подстроки
strings.Contains("hello", "ell")
Обработка ошибок
Всегда проверяйте ошибки при использовании regexp.Compile (не MustCompile):

go
re, err := regexp.Compile(`complex(\d+)pattern`)
if err != nil {
    log.Fatal("Invalid regex pattern:", err)
}
Продвинутые техники
1. Жадные и ленивые квантификаторы
go
re := regexp.MustCompile(`a.*b`)   // Жадный: "aabab" → "aabab"
re := regexp.MustCompile(`a.*?b`)  // Ленивый: "aabab" → "aab"
2. Многострочный режим
go
re := regexp.MustCompile(`(?m)^start`) // ^ соответствует началу каждой строки
3. Флаги модификаторы
go
re := regexp.MustCompile(`(?i)hello`) // Регистронезависимый поиск
*/
