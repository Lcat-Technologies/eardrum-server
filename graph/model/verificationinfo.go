package model

import (
	"github.com/GigaDesk/eardrum-prefix/validate"
)

// validates Verificationinfo input data
func (n Verificationinfo) Validate() error {
	//validate phone number in kenyan conditions
	if err := validate.ValidateKenyanPhoneNumber(n.PhoneNumber); err != nil {
		return err
	}
	//validate otp
    if err := validate.ValidateOtp(n.Otp); err != nil {
		return err
	}
	return nil

}
