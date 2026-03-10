package src

import (
	"ShoppingCalculator/helper"
	// "fmt"
	// "strconv"
)

// this func is supposed to take all the output from the input func and then sum up all the total, the total onumber of items and the total sumxx amount

func Calc(k []helper.Input) helper.Output{
	var total helper.Output 
	// result := []helper.Output{}
	num := 0
	var mun float64 = 0
	for _, c := range k{
		num += c.NumberOfItems
		total.TotalQuantity = num
		mun += float64(c.Cost)
		total.TotalCost = mun
	}
	return  total
}