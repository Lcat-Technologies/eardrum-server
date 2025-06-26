package model

import (
	"github.com/GigaDesk/eardrum-prefix/validate"
)

// validates NewShop input data
func (n NewShop) Validate() error {

	//validate name
	if err := validate.ValidateName(n.Name); err != nil {
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

// returns the shop's name
func (n NewShop) GetName() string {
	return n.Name
}

// returns the shop's phone number
func (n NewShop) GetPhoneNumber() string {
	return n.PhoneNumber
}

// returns the shop's password
func (n NewShop) GetPassword() string {
	return n.Password
}