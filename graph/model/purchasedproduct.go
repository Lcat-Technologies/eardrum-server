package model

// returns the purchased product id
func (p PurchasedProduct) GetProductID() int64 {
	return int64(p.ProductID)
}

// returns the number of units of the product bought
func (p PurchasedProduct) GetUnitsBought() int {
	return  p.UnitsBought
}
