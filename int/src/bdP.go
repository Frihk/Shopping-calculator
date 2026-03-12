package src

import (
	"encoding/json"
	"os"

	// "encoding"
	"ShoppingCalculator/helper"
	// "ShoppingCalculator/int/src"
)

// this script is incharge of saving, editing the data, loading the data from the json file where data is saved.
// I'll also try to make sure the data is sorted by the frequency/ how many times the consumer buys the commodity, this will help when am making the suggestions

func Jupdate() error {
	filepath := "/home/fian/frihk/ShoppingCalculator/storage/shopinglogs.json"
	_, newProduct := Input()
	var products []helper.ProductStorage
	// open the file
	file, err := os.Open(filepath)
	if err == nil {
		defer file.Close()
		json.NewDecoder(file).Decode(&products)
	}
	for _, c := range newProduct{
		found := false
		for i := range products {
			if products[i].Name == c.Name {
				products[i].Freq++
				found = true
				break
			}
		}
		if !found {
			c.Freq = 1
			products = append(products, c)
		}
	}

	file, err = os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(products)
	if err != nil {
		return err
	}
	return nil
}
