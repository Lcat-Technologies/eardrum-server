package encrypt

import (
	"github.com/Lcat-Technologies/eardrum-interfaces/errors"
	"golang.org/x/crypto/bcrypt"
)

// CheckPassword compares a provided password with the hashed version to find out if they match. It returns nil on success or error if they don't match
func CheckPassword(hashedpassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedpassword), []byte(password))
	if err != nil {
		err1 := errors.New(errors.EARAuthPasswordMismatch, err)
		err1.Log()
		return err1
	}
	return nil
}
