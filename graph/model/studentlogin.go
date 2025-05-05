package model

import "github.com/GigaDesk/eardrum-prefix/validate"

// validates student login input data
func (n StudentLogin) Validate() error {

	//validate school id
	if err := validate.ValidateId(n.Schoolid); err != nil {
		return err
	}
	//validate registration number
	if err := validate.ValidateRegistrationNumber(n.RegistrationNumber); err != nil {
		return err
	}

	//validate password
	if err := validate.ValidatePassword(n.Password); err != nil {
		return err
	}
	return nil
}
