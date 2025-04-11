package main

import (
	"fmt"
	"strings"
)

const USD_TO_RUB float64 = 83.74
const USD_TO_EUR float64 = 1.137

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
	var usdValue float64
	switch from {
	case "USD":
		usdValue = value
	case "RUB":
		usdValue = value / USD_TO_RUB
	case "EUR":
		usdValue = value / USD_TO_EUR
	default:
		return 0
	}

	switch to {
	case "USD":
		return usdValue
	case "RUB":
		return usdValue * USD_TO_RUB
	case "EUR":
		return usdValue * USD_TO_EUR
	default:
		return 0
	}
}
