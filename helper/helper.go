package helper

type Input struct {
	ItemName 		string
	NumberOfItems 	int
	PriceOfItem 	float64
	Cost 			float64

}

type Output struct {
	TotalQuantity 	int
	TotalCost 		float64
}

type ProductStorage struct{
	Name 	string `json:"name"`
	Price 	float64	`json:"price"`
	Freq 	int		`json:"freq"`
}

