package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	operations := []string{"AVG", "SUM", "MED"}
	for {
		operation := scanInput("Введите операцию (AVG - среднее, SUM - сумма, MED - медиана): ")
		if !slices.Contains(operations, operation) {
			fmt.Println("Неверная операция, попробуйте снова")
			continue
		}

		numbers := scanNumbers("Введите числа через запятую: ")
		if len(numbers) == 0 {
			fmt.Println("Ни одно число не было введено")
			continue
		}

		result := calc(operation, numbers)
		fmt.Printf("Результат: %v\n\n", result)
	}
}

func scanInput(message string) string {
	fmt.Println(message)
	var input string
	fmt.Scanln(&input)
	return input
}

func scanNumbers(message string) []int {
	fmt.Println(message)
	var input string
	fmt.Scanln(&input)

	parts := strings.Split(input, ",")
	var numbers []int
	for _, p := range parts {
		num, err := strconv.Atoi(strings.TrimSpace(p))
		if err != nil {
			fmt.Printf("Некорректное значение: %q\n", p)
			continue
		}
		numbers = append(numbers, num)
	}
	return numbers
}

func calc(operation string, numbers []int) int {
	switch operation {
	case "AVG":
		return calcAvg(numbers)
	case "SUM":
		return calcSum(numbers)
	case "MED":
		return calcMed(numbers)
	default:
		return 0
	}
}

func calcAvg(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum / len(numbers)
}

func calcSum(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func calcMed(numbers []int) int {
	slices.Sort(numbers)
	mid := len(numbers) / 2
	if len(numbers)%2 == 0 {
		return (numbers[mid-1] + numbers[mid]) / 2
	}
	return numbers[mid]
}
