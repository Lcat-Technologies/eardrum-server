package model

// returns the product's name
func (n NewProduct) GetName() string {
	return n.Name
}

// returns the product's price per unit in cents
func (n NewProduct) GetPricePerUnitInCents() int64 {
	return int64(n.PricePerUnitInCents)
}