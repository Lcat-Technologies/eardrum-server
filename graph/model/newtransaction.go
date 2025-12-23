package model

import (
	"github.com/GigaDesk/eardrum-interfaces/transaction"
)

// returns the transaction's universal unique identifier
func (n NewProductTransaction) GetUUID() string {
	return n.QRCode
}

// returns the transaction's pin code
func (n NewProductTransaction) GetPinCode() string {
	return n.PinCode
}

// returns the transaction's purchased products
func (n NewProductTransaction) GetPurchasedProducts() []transaction.PurchasedProduct {

	var list []transaction.PurchasedProduct
	for _, purchasedproduct := range n.PurchasedProducts {
		list = append(list, purchasedproduct)
	}
	return list
}

// returns the transaction's universal unique identifier
func (n NewAmountTransaction) GetUUID() string {
	return n.QRCode
}

// returns the transaction's pin code
func (n NewAmountTransaction) GetPinCode() string {
	return n.PinCode
}

// returns the transaction's amount
func (n NewAmountTransaction) GetTotalAmountInCents() uint {
	return uint(n.AmountInCents)
}
