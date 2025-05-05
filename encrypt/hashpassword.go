package encrypt

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"github.com/rs/zerolog/log"
)

//encrypts plain text passwords
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil { 
		log.Error().Str("password", password).Msg(err.Error())
		return "", errors.New("error occurred during password encryption")
	}
	return string(bytes), nil
}