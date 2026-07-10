package model

//Returns the facial embedding to authorize transaction
func (n NewAmountTransaction) GetFacialEmbedding() string {
return n.FacialEmbedding
}

//Returns the amount to be processed in the transaction
func (n NewAmountTransaction) GetTotalAmountInCents() uint{
return uint(n.AmountInCents)
}

//Returns the unique identifier to the user account
func (n NewAmountTransaction) GetUUID() string{
return n.QRCode
}