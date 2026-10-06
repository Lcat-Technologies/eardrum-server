package model

import (
"time"
"github.com/Lcat-Technologies/eardrum-prefix/validate"
)

//Returns the facial embedding to authorize transaction
func (n NewOnlineTransaction) GetFacialEmbedding() string {
return n.FacialEmbedding
}

//Returns the amount to be processed in the transaction
func (n NewOnlineTransaction) GetTotalAmountInCents() uint{
return uint(n.AmountInCents)
}

//Returns the unique identifier to the user account
func (n NewOnlineTransaction) GetUUID() string{
return n.QRCode
}

//OFFLINE TRANSACTIONS


//Returns the facial embedding to authorize transaction
func (n NewOfflineTransaction) GetFacialEmbedding() string {
return n.FacialEmbedding
}

//Returns the amount to be processed in the transaction
func (n NewOfflineTransaction) GetTotalAmountInCents() uint{
return uint(n.AmountInCents)
}

//Returns the unique identifier to the user account
func (n NewOfflineTransaction) GetUUID() string{
return n.QRCode
}

//Retuns the offline timestamp of the transaction
func (n NewOfflineTransaction) GetOfflineTimestamp() time.Time{
	return n.OfflineTimeStamp
}


//Retuns the user's phone number
func (n NewOfflineTransaction) GetPhoneNumber() string{
	return n.PhoneNumber
}

//Return's base64 encoding of the image of the scan that authorized transaction
func (n NewOfflineTransaction) GetScanLog() string{
	return n.ScanLog
}

//Returnd the OfflineTransactionID
func (n NewOfflineTransaction) GetOfflineTransactionID() string{
	return n.OfflineTransactionID
}


//Validate NewOfflineTransaction input data
func(n NewOfflineTransaction) Validate() error {
	//validate offline transaction id
	if err := validate.ValidateOfflineTransactionID(n.OfflineTransactionID); err != nil {
		return err
	}
	return nil
}

