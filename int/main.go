package main

import (
	"fmt"
	"ShoppingCalculator/int/src"
	
)

var _ fmt.Formatter
// type Input struct {
// 	ItemName string
// 	NumberOfItems string 
// 	PriceOfItem string
// 	Cost string

// }
  
	 
func main() {
	list := src.Input()
	fmt.Println(src.Calc(list))
}