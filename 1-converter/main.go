package main

import (
	"fmt"
	"strings"
)

var currencyMap = map[string]float64{
	"USD": 82.17,
	"EUR": 93.45,
	"RUB": 1.0,
}

func main() {
	for {
		currencies := "USD EUR RUB"
		currentCurrency := ""
		for {
			currentCurrency = scanCurrency("Введите текущую валюту: " + currencies)
			if !strings.Contains(currencies, currentCurrency) || currentCurrency == "" {
				continue
			}
			break
		}

		value := 0.0
		for {
			value = scanValue("Введите количество")
			if value <= 0 {
				continue
			}
			break
		}

		targetCurrency := ""
		for {
			availableCurrency := strings.ReplaceAll(currencies, currentCurrency, "")
			targetCurrency = scanCurrency("Введите желаемую валюту: " + availableCurrency)
			if !strings.Contains(availableCurrency, targetCurrency) || targetCurrency == "" {
				continue
			}
			break
		}

		result := convert(value, currentCurrency, targetCurrency)
		fmt.Printf("%.2f %s = %.2f %s\n\n", value, currentCurrency, result, targetCurrency)
		return
	}
}

func scanCurrency(message string) (value string) {
	fmt.Println(message)
	fmt.Scanln(&value)
	return value
}

func scanValue(message string) (value float64) {
	fmt.Println(message)
	fmt.Scanln(&value)
	return value
}

func convert(value float64, from string, to string) float64 {
	rubles := value * currencyMap[from]
	return rubles / currencyMap[to]
}
