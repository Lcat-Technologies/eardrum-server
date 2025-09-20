package model

import "github.com/GigaDesk/eardrum-interfaces/transaction"

// returns the transaction's phone number
func (n NewTransaction) GetPhoneNumber() string {
	return n.PhoneNumber
}

// returns the transaction's pin code
func (n NewTransaction) GetPinCode() string {
	return n.PinCode
}

//returns the transaction's purchased products
func (n NewTransaction) GetPurchasedProducts() []transaction.PurchasedProduct{

	var list []transaction.PurchasedProduct
    for _ , purchasedproduct:= range n.PurchasedProducts{
    list = append(list, purchasedproduct)
	}
	return list
}