package src

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sort"

	// "encoding"
	"ShoppingCalculator/helper"
	// "ShoppingCalculator/int/src"
)

// this script is incharge of saving, editing the data, loading the data from the json file where data is saved.
// I'll also try to make sure the data is sorted by the frequency/ how many times the consumer buys the commodity, this will help when am making the suggestions

func logPath() string {
	if override := os.Getenv("SHOPPING_LOG_PATH"); override != "" {
		return override
	}
	return filepath.Join("storage", "shopinglogs.json")
}

func Jupdate(newProduct []helper.ProductStorage) error {
	logFilePath := logPath()
	// _, newProduct := Input()
	var products []helper.ProductStorage
	// open the file
	file, err := os.Open(logFilePath)
	if err == nil {
		defer file.Close()
		json.NewDecoder(file).Decode(&products)
	}
	for _, c := range newProduct {
		found := false
		for i := range products {
			if products[i].Name == c.Name {
				products[i].Freq++
				products[i].Price = c.Price
				found = true
				break
			}
		}
		if !found {
			c.Freq = 1
			products = append(products, c)
		}
	}

	file, err = os.Create(logFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	sort.Slice(products, func(i, j int) bool {
		return products[i].Freq > products[j].Freq
	})

	encode := json.NewEncoder(file)
	encode.SetIndent("", "  ")
	err = encode.Encode(products)
	if err != nil {
		return err
	}
	return nil
}
