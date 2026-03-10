package src

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	// "ShoppingCalculator/helper"
	"ShoppingCalculator/helper"
)


func Input( ) []helper.Input {
	var input helper.Input
	k := []helper.Input{}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		// if scanner.Scan() {
		fmt.Print("name of the item: ")
		scanner.Scan()
		input.ItemName = scanner.Text()
		if input.ItemName == "" {
			break
		}

		fmt.Print("Quantiy of the item: ")
		scanner.Scan()
		quantity :=scanner.Text()
		num, err := strconv.Atoi(quantity)
		if err != nil{
			fmt.Println("Invalid Quantity")
		}
		input.NumberOfItems = num

		fmt.Print("Price of the item: ")
		scanner.Scan()
		price := scanner.Text()
		cash, err := strconv.ParseFloat(price, 32)
			if err != nil {
			fmt.Println("Invalid Price")
		}
		input.PriceOfItem = cash
		val := float64(num) * cash
		input.Cost =  float64(val)//strconv.FormatFloat(val, 'f', 2, 32) 
		// }
		k = append(k, input)
		
	}
		// int := Input(input.ItemName, input.NumberOfItems, input.PriceOfItem, input.Cost)
		return k

}

