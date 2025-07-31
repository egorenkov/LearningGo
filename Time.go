package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func howOldAreYou(start, end time.Time) {
	diff := end.Sub(start)
	years := diff.Hours() / (24 * 365) // Более точный расчёт
	fmt.Printf("Сейчас вам %d лет\n", int(years))
}

func timer(hour, minute, second int) {
	fmt.Printf("Запуск таймера на %d ч %d мин %d сек\n", hour, minute, second)
	totalSeconds := hour*3600 + minute*60 + second
	time.Sleep(time.Duration(totalSeconds) * time.Second)
	fmt.Printf("Таймер сработал: %d ч %d мин %d сек\n", hour, minute, second)
}

func main() {

	fmt.Println(`Добро пожаловать в консольное приложение!
	Доступные команды:
	- help (для просмотра доступных команд)
	- timer (для работы с таймером)
	- birthday (для того, чтобы узнать сколько прошло времени с дня рождения)
	- sec (для работы с секундомером)
	- exit (для выхода)`)
	reader := bufio.NewReader(os.Stdin)

	var numForTime int

	var start time.Time
	var end time.Time
	flagForTime := false
	flagForSecondTouch := false
	for {

		input, _ := reader.ReadString('\n')

		strForWork := strings.TrimSpace(input)

		if strForWork == "sec" {
			numForTime++
			flagForTime = true
			flagForSecondTouch = true
		} else if strForWork == "timer" {
			fmt.Println("Введите время для засечения в формате: *hour* *minute* *seconds*")
			var hour int
			var min int
			var sec int
			fmt.Scanf("%d %d %d", &hour, &min, &sec)
			timer(hour, min, sec)
		} else if strForWork == "birthday" {
			fmt.Println("Введите дату рождения в формате *year* *mounth* *day*")
			var year int
			var mont int
			var day int
			fmt.Scanf("%d %d %d", &year, &mont, &day)
			if mont >= 1 && mont <= 12 && day >= 1 && day <= 31 {
				birt := time.Date(year, time.Month(mont), day, 0, 0, 0, 0, time.UTC)
				howOldAreYou(birt, time.Now())
			} else {
				fmt.Println(mont)
				fmt.Println("Некорректный месяц или день")
				fmt.Println(`        Доступные команды:
	- help (для просмотра доступных команд)
	- timer (для работы с таймером)
	- birthday (для того, чтобы узнать сколько прошло времени с дня рождения)
	- sec (для работы с секундомером)
	- exit (для выхода) `)
			}

		} else if strForWork == "exit" {
			break
		} else if strForWork == "help" {
			fmt.Println(`        Доступные команды:
	- help (для просмотра доступных команд)
	- timer (для работы с таймером)
	- birthday (для того, чтобы узнать сколько прошло времени с дня рождения)
	- sec (для работы с секундомером)
	- exit (для выхода) `)
		} else if strForWork == "" && flagForSecondTouch {
			numForTime++
			flagForSecondTouch = false

		} else if strForWork == "" {

		} else {
			fmt.Println("Некорректный ввод, попробуйте снова")
		}

		//For secundomer
		if numForTime == 1 && flagForTime {
			fmt.Println("Секундомер запущен")
			start = time.Now()
			fmt.Println("Нажмите enter чтобы остановить его")
		} else if numForTime == 2 && flagForTime {
			end = time.Now()
			flagForTime = false
			fmt.Printf("Секундомер отработал на %f секунд", end.Sub(start).Seconds())
			numForTime = 0
		}

	}

	fmt.Println("Спасибо за пользование консольное приложение!")

}
