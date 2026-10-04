package main

import "fmt"

func main() {

	// Задание 1
	temperature := 15

	if temperature < 0 {
		fmt.Println("Холодно")
	} else if temperature <= 20 {
		fmt.Println("Тепло")
	} else {
		fmt.Println("Жарко")
	}

	// Задание 2
	score := 85

	if score >= 90 {
		fmt.Println("Отлично")
	} else if score >= 70 {
		fmt.Println("Хорошо")
	} else if score >= 50 {
		fmt.Println("Удовлетворительно")
	} else {
		fmt.Println("Не сдал")
	}

	// Задание 3
	hour := 14

	switch {
	case hour >= 0 && hour <= 5:
		fmt.Println("Ночь")
	case hour >= 6 && hour <= 11:
		fmt.Println("Утро")
	case hour >= 12 && hour <= 17:
		fmt.Println("День")
	case hour >= 18 && hour <= 23:
		fmt.Println("Вечер")
	}

	// Задание 4
	var number int
	fmt.Print("Введите число: ")
	fmt.Scan(&number)

	if number%2 == 0 {
		fmt.Println("Чётное число")
	} else {
		fmt.Println("Нечётное число")
	}

	// Задание 5
	day := "Monday"

	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("Будний день")
	case "Saturday", "Sunday":
		fmt.Println("Выходной")
	default:
		fmt.Println("Некорректный день")
	}

	// Задание 6
	balance := 100

	if balance >= 0 {
		fmt.Println("Баланс положительный")
	} else {
		fmt.Println("Баланс отрицательный")
	}

	// Задание 7
	var age int
	fmt.Print("Введите возраст: ")
	fmt.Scan(&age)

	if age < 13 {
		fmt.Println("Ребёнок")
	} else if age <= 17 {
		fmt.Println("Подросток")
	} else {
		fmt.Println("Взрослый")
	}

	// Задание 8
	var command string
	fmt.Print("Введите команду: ")
	fmt.Scan(&command)

	switch command {
	case "start":
		fmt.Println("Запуск")
	case "stop":
		fmt.Println("Остановка")
	case "restart":
		fmt.Println("Перезапуск")
	default:
		fmt.Println("Неизвестная команда")
	}

	// Задание 9
	grade := 3

	switch grade {
	case 1:
		fmt.Println("A")
	case 2:
		fmt.Println("B")
	case 3:
		fmt.Println("C")
	case 4:
		fmt.Println("D")
	case 5:
		fmt.Println("F")
	default:
		fmt.Println("Некорректная оценка")
	}
}