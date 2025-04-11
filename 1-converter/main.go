package main

import "fmt"

const USD_TO_RUB float64 = 83.74
const USD_TO_EUR float64 = 1.137

func main() {
	eur := 2.0
	res := EurToRub(eur)
	fmt.Println(res)
}

func EurToRub(value float64) float64 {
	usd := value * USD_TO_EUR
	return usd * USD_TO_RUB
}
