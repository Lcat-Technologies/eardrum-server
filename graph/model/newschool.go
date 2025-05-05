package model

import (
	"github.com/AlekSi/pointer"
	"github.com/GigaDesk/eardrum-prefix/validate"
)

// validates Newschool input data
func (n NewSchool) Validate() error {
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

	//validate badge
	if n.Badge != nil {
		if err := validate.ValidateBadge(pointer.GetString(n.Badge)); err != nil {
			return err
		}
	}

	//validate website
	if n.Website != nil {
		if err := validate.ValidateWebsite(pointer.GetString(n.Website)); err != nil {
			return err
		}
	}
	return nil
}
