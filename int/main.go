package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"ShoppingCalculator/helper"
	"ShoppingCalculator/int/src"
	// "int/src"
	// "Shopping-calculator/int/src"
)


// type Input struct {
// 	ItemName string
// 	NumberOfItems string 
// 	PriceOfItem string
// 	Cost string

// }
  
	 
func main() {
	var input helper.Input
	// ItemName.item := ""
	// quantity := ""
	// price := ""

	scanner := bufio.NewScanner(os.Stdin)
	// if scanner.Scan() {
	fmt.Print("name of the item: ")
	scanner.Scan()
	input.ItemName = scanner.Text()

	fmt.Print("Quantiy of the item: ")
	scanner.Scan()
	quantity :=scanner.Text()
	num, err := strconv.ParseFloat(quantity, 64)
	if err != nil{
		fmt.Println("Invalid Quantity")
	}
	input.NumberOfItems = quantity

	fmt.Print("Price of the item: ")
	scanner.Scan()
	price := scanner.Text()
	cash, err := strconv.ParseFloat(price, 64)
		if err != nil {
		fmt.Println("Invalid Price")
	}
	input.PriceOfItem = price
	val := num * cash
	input.Cost = strconv.FormatFloat(val, 'f', 2, 64) 
	// }
	k := src.Input(input)
	// int := Input(input.ItemName, input.NumberOfItems, input.PriceOfItem, input.Cost)
	fmt.Println(k[0])
}