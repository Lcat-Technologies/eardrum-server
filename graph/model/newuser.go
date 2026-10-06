package model

import (
	"github.com/Lcat-Technologies/eardrum-prefix/validate"
	"github.com/google/uuid"
)

// validates NewUser input data
func (n NewUser) Validate() error {

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

// returns the user's name
func (n NewUser) GetUserName() string {
	return n.Username
}

// returns the user's phone number
func (n NewUser) GetPhoneNumber() string {
	return n.PhoneNumber
}

// returns the user's password
func (n NewUser) GetPassword() string {
	return n.Password
}


// ParseUUIDSlice converts a slice of string UUIDs to a slice of uuid.UUID structs.
func ParseUUIDSlice(strs []string) ([]uuid.UUID) {
	uuids := make([]uuid.UUID, len(strs))
	for i, str := range strs {
		parsed, _ := uuid.Parse(str)
		uuids[i] = parsed
	}
	return uuids
}
