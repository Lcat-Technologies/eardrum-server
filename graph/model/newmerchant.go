package model

import (
	"github.com/Lcat-Technologies/eardrum-prefix/validate"
)

// validates NewMerchant input data
func (n NewMerchant) Validate() error {

	//validate name
	if err := validate.ValidateName(n.Username); err != nil {
		return err
	}

	//validate phone number in kenyan conditions
	if err := validate.ValidateKenyanPhoneNumber(n.PhoneNumber); err != nil {
		return err
	}


	//validate password
	if err := validate.ValidatePassword(n.Password); err != nil {
		return err
	}

	return nil
}

// returns the merchant's name
func (n NewMerchant) GetUserName() string {
	return n.Username
}

// returns the merchant's phone number
func (n NewMerchant) GetPhoneNumber() string {
	return n.PhoneNumber
}

// returns the merchant's password
func (n NewMerchant) GetPassword() string {
	return n.Password
}


