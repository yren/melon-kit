package main

import (
	"fmt"
	"github.com/shopspring/decimal"
)

func main() {
	price, err := decimal.NewFromString("136.02")
	if err != nil {
		panic(err)
	}

	quantity := decimal.NewFromFloat(3)
	fee, _ := decimal.NewFromString(".035")
	taxRate, _ := decimal.NewFromString(".08875")

	subtotal := price.Mul(quantity)
	preTax := subtotal.Mul(fee.Add(decimal.NewFromFloat(1)))
	total := preTax.Mul(taxRate.Add(decimal.NewFromFloat(1)))
	fmt.Println("Subtotal:", subtotal)
	fmt.Println("PreTax:", preTax)
	fmt.Println("Taxes:", total.Sub(preTax))
	fmt.Println("Total:", total)
	fmt.Println("Tax rate:", total.Sub(preTax).Div(preTax))

	}
