package model

import (
	"github.com/GigaDesk/eardrum-prefix/validate"
)

// validates NewUser input data
func (n NewUser) Validate() error {

	//validate name
	if err := validate.ValidateName(n.Name); err != nil {
		return err
	}

	//validate phone number in kenyan conditions
	if err := validate.ValidateKenyanPhoneNumber(n.PhoneNumber); err != nil {
		return err
	}

	//validate phone number in kenyan conditions
	if err := validate.ValidateKenyanPhoneNumber(n.MpesaNumber); err != nil {
		return err
	}

	//validate password
	if err := validate.ValidatePassword(n.Password); err != nil {
		return err
	}

	return nil
}

// returns the user's name
func (n NewUser) GetName() string {
	return n.Name
}

// returns the user's phone number
func (n NewUser) GetPhoneNumber() string {
	return n.PhoneNumber
}

// returns the user's password
func (n NewUser) GetPassword() string {
	return n.Password
}

// returns the user's mpesa number
func (n NewUser) GetMpesaNumber() string {
	return n.MpesaNumber
}
