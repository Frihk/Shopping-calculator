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
	inputs,_ := src.Input()
	fmt.Println(src.Calc(inputs))
}