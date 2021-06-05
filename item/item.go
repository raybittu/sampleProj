package item

type Item struct {
	Category string
	Brand string
	Price float64
}

func Create(category, brand string, price float64) *Item {
	return &Item{Category: category, Brand:brand, Price:price}
}