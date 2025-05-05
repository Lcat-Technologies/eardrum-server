package model

import (
	"github.com/AlekSi/pointer"
	"github.com/GigaDesk/eardrum-prefix/validate"
)

// validates Newstudent input data
func (n NewStudent) Validate() error {

	//validate registration number
	if err := validate.ValidateRegistrationNumber(n.RegistrationNumber); err != nil {
		return err
	}

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

	//validate profile picture
	if n.ProfilePicture != nil {
		if err := validate.ValidateProfilePicture(pointer.GetString(n.ProfilePicture)); err != nil {
			return err
		}
	}

	//validate date of admission
	if n.DateOfAdmission != nil {
		if err := validate.ValidateDateOfAdmission(pointer.GetTime(n.DateOfAdmission)); err != nil {
			return err
		}
	}
    

	//validate date of birth
	if n.DateOfBirth != nil {
		if err := validate.ValidateDateOfBirth(pointer.GetTime(n.DateOfBirth)); err != nil {
			return err
		}
	}

	return nil
}
