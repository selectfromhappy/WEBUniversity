package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	//Пункт 1
	veryBadUnclearName := "15 chicken wings"
	order := &veryBadUnclearName
	newOrder := *order + " " + " Wendy's "
	fmt.Println(newOrder)

	//Пункт 2 // Числа
	numberA := 800
	fmt.Print("Целочисленное значение 1: ", numberA)
	numberB := 900
	fmt.Print(" Целочисленное значение 2 : ", "\n", numberB)
	numberC := 900.3
	fmt.Println("\n", "Дробное значение:", numberC)
	fmt.Println(8 + 4)
	var lastMonth, thisMonth int
	lastMonth = 1000
	thisMonth = 3000
	fmt.Println("Разница между расходами: ", thisMonth-lastMonth)

	//Пункт 3 // Умножение и деление
	var numLanguages, months, days int
	numLanguages = 4
	months = 11
	days = months * 16
	var daysPerLanguage float64 = float64(days) / float64(numLanguages)
	fmt.Println("Дни на языки: ", daysPerLanguage)

	//Пункт 4 //Степень
	fmt.Println("Оператора ** в языке Golang не существует")
	fmt.Println(math.Pow(26.0, 2.0))

	//Пункт 5 //Операторы присвоения
	myNum := 5
	answer := myNum
	fmt.Println("Проверяем присвоение к answer: ", answer)
	answer += 2
	fmt.Println("Проверяем answer +=2 : ", answer)
	answer *= 2
	fmt.Println("Проверяем answer *=2 : ", answer)
	answer -= 2
	fmt.Println("Проверяем answer -=2 : ", answer)
	answer /= 2
	fmt.Println("Проверяем answer /=2 : ", answer)
	fmt.Println("Разность исходного числа и answer: ", myNum-answer)

	//Пункт 6 //Математические Функции
	var a, b int
	a = 10
	b = 3
	fmt.Println(a % b)
	if a%b == 0 {
		fmt.Println("Делится", a/b)
	} else {
		fmt.Println("Делится с остатком: ", a%b)
	}
	numberPow := 2.0
	st := math.Pow(numberPow, 10.0)
	fmt.Println("Число в 10-й степени: ", st)
	fmt.Println("Корень из 245 = ", math.Sqrt(245))
	var numbers [7]float64 = [7]float64{4, 2, 5, 19, 13, 0, 10}
	var sum float64
	for _, number := range numbers {
		sum += number * number
	}
	fmt.Println("Корень из суммы квадратов элементов массива: ", math.Sqrt(sum))
	floorNumber := math.Sqrt(379)
	fmt.Printf("Квадратный корень из 379: \nЦелое: %d\nДесятые: %.1f\nСотые: %.2f\n",
		int(math.Round(floorNumber)),
		math.Round(floorNumber*10)/10,
		math.Round(floorNumber*100)/100)

	var numbersMinMax [7]float64 = [7]float64{4, -2, 5, 19, -130, 0, 10}
	minNumber, maxNumber := math.MaxFloat64, -math.MaxFloat64
	for _, num := range numbersMinMax {
		if num < minNumber {
			minNumber = num
		}
		if num > maxNumber {
			maxNumber = num
		}
	}
	fmt.Printf("Минимальное: %.2f, Максимальное: %.2f\n", minNumber, maxNumber)
	fmt.Println("Случайное число от 0 до 100", rand.Intn(101))

	var numbersRandomArray [9]int
	for i, _ := range numbersRandomArray {
		numberGen := rand.Int()
		numbersRandomArray[i] = numberGen
	}
	for i, n := range numbersRandomArray {
		fmt.Println(i, n)
	}
	numbA := 15
	numbB := 900
	fmt.Println(math.Abs(float64(numbA) - float64(numbB)))

	var positiveNegativeNumber [6]int = [6]int{1, 2, -1, -2, 3, -3}
	fmt.Println("До: ")
	for i, number := range positiveNegativeNumber {
		fmt.Println(i, number)
	}
	for i, number := range positiveNegativeNumber {
		if number < 0 {
			positiveNegativeNumber[i] = positiveNegativeNumber[i] * (-1)
		}
	}
	fmt.Println("После: ")
	for i, number := range positiveNegativeNumber {
		fmt.Println(i, number)
	}

	numberDel := 30 //Можно через Scan делать, чтобы вводить числа
	fmt.Println("Делители числа: ")
	var numberDelNumbers []int
	for i := 1; i <= numberDel; i++ {
		if numberDel%i == 0 {
			numberDelNumbers = append(numberDelNumbers, i)
		}
	}
	fmt.Println(numberDelNumbers)

	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sumElementsArray := 0
	count := 0

	// Проход по элементам массива
	for _, value := range array {
		sumElementsArray += value
		count++

		if sumElementsArray > 10 {
			break
		}
	}
	fmt.Printf("Количество первых элементов, которые нужно сложить: %d\n", count)

	str := "Sveiki" // Пример строки
	myNum = printStringReturnNumber(str)
	fmt.Printf("Длина строки: %d\n", myNum)
}

func printStringReturnNumber(s string) int {
	fmt.Println(s)
	return len(s)
}
