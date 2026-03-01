package src

// import (
// 	"ShoppingCalculator/helper"
// )


// func Input(s helper.Input) []helper.Input {
// 	var input helper.Input
// 	k := []helper.Input{}

// 	scanner := bufio.NewScanner(os.Stdin)

// 	for {
// 		// if scanner.Scan() {
// 		fmt.Print("name of the item: ")
// 		scanner.Scan()
// 		input.ItemName = scanner.Text()
// 		if input.ItemName == "" {
// 			break
// 		}

// 		fmt.Print("Quantiy of the item: ")
// 		scanner.Scan()
// 		quantity :=scanner.Text()
// 		num, err := strconv.ParseFloat(quantity, 64)
// 		if err != nil{
// 			fmt.Println("Invalid Quantity")
// 		}
// 		input.NumberOfItems = quantity

// 		fmt.Print("Price of the item: ")
// 		scanner.Scan()
// 		price := scanner.Text()
// 		cash, err := strconv.ParseFloat(price, 64)
// 			if err != nil {
// 			fmt.Println("Invalid Price")
// 		}
// 		input.PriceOfItem = price
// 		val := num * cash
// 		input.Cost = strconv.FormatFloat(val, 'f', 2, 64) 
// 		// }
// 		k = append(k, input)
		
// 	}
// 		// int := Input(input.ItemName, input.NumberOfItems, input.PriceOfItem, input.Cost)
// 		fmt.Println(k)

// }

