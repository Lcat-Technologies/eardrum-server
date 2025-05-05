package encrypt

import "golang.org/x/crypto/bcrypt"


//CheckPassword compares a provided password with the hashed version to find out if they match. It returns nil on success or error if they don't match
func CheckPassword(hashedpassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedpassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}