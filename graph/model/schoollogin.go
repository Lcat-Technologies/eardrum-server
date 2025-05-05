package model

import (
	"github.com/GigaDesk/eardrum-prefix/validate"
)

// validates school login input data
func (n SchoolLogin) Validate() error {
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