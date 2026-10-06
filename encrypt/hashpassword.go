package encrypt

import (

	"golang.org/x/crypto/bcrypt"
	"github.com/rs/zerolog/log"
	"github.com/Lcat-Technologies/eardrum-interfaces/errors"
)

//encrypts plain text passwords
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil { 
		log.Error().Str("password", password).Msg(err.Error())
		err1 := errors.New(errors.EARAuthPasswordHashingFailed, err)
		err1.Log()
		return "", err1
	}
	return string(bytes), nil
}