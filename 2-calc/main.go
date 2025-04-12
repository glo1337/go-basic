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
			continue
		}

		values := scanInput("Введите числа через запятую: ")
		valuesSplitted := strings.Split(values, ",")
		if len(values) == 0 {
			continue
		}
		res := calc(operation, valuesSplitted)
		resString := fmt.Sprintf("Результат: %v", res)
		fmt.Println(resString)
	}
}

func scanInput(message string) (value string) {
	fmt.Println(message)
	fmt.Scanln(&value)
	return value
}

func calc(operation string, values []string) int {
	res := 0
	switch operation {
	case "AVG":
		res = calcAvg(values)
	case "SUM":
		res = calcSum(values)
	case "MED":
		res = calcMed(values)
	}
	return res
}

func calcAvg(values []string) int {
	res := 0
	for _, v := range values {
		number, err := strconv.Atoi(strings.TrimSpace(v))
		if err != nil {
			continue
		}
		res += number
	}
	return res / len(values)
}

func calcSum(values []string) int {
	res := 0
	for _, v := range values {
		number, err := strconv.Atoi(strings.TrimSpace(v))
		if err != nil {
			continue
		}
		res += number
	}
	return res
}

func calcMed(values []string) int {
	var numbers []int
	for _, v := range values {
		number, err := strconv.Atoi(strings.TrimSpace(v))
		if err != nil {
			continue
		}
		numbers = append(numbers, number)
	}

	if len(numbers) == 0 {
		return 0
	}

	slices.Sort(numbers)

	mid := len(numbers) / 2
	if len(numbers)%2 == 0 {
		return (numbers[mid-1] + numbers[mid]) / 2
	}
	return numbers[mid]
}
