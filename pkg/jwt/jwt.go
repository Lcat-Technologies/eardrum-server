package jwt

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/rs/zerolog/log"
	"github.com/GigaDesk/eardrum-interfaces/errors"
)

// secret key being used to sign tokens
var (
	JwtSecretKey []byte
)

type TokenCredentials struct {
    Username   string
	Role string
}

// returns the role of a given user
func (tokencredentials *TokenCredentials) GetRole() string {
	return tokencredentials.Role
}

// returns the username of a given user/merchant
func (tokencredentials *TokenCredentials) GetUsername() (string) {
	return tokencredentials.Username
}

// GenerateToken takes token credentials generates a jwt token and assign an id and role to it's claims and return it. It returns an error on failure
func GenerateToken(credentials TokenCredentials) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	/* Create a map to store our claims */
	claims := token.Claims.(jwt.MapClaims)
	/* Set token claims */
	claims["username"] = credentials.Username
	claims["role"] = credentials.Role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString(JwtSecretKey)
	if err != nil {
		err1 := errors.New(errors.EARJwtSigningFailed, err)
		err1.Log()
		return "", err1
	}
	return tokenString, nil
}

// ParseToken parses a jwt token and returns the id and role in it's claims as the token's credentials
func ParseToken(tokenStr string) (*TokenCredentials, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return JwtSecretKey, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		credentials := &TokenCredentials{
			Username:   claims["username"].(string),
			Role: claims["role"].(string),
		}
		return credentials, nil
	} else {
		err1 := errors.New(errors.EARJwtParsingFailed, err)
		err1.Log()
		return nil, err1
	}
}

// Initializes the JwtSecretKey global variable
func InitializeJwtSecretKey() {
	JwtSecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))
	log.Trace().Str("secretkey", string(JwtSecretKey)).Msg("completed jwt secret key initialization")
}
