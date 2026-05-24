package security

import "golang.org/x/crypto/bcrypt"

//The Hash function receive a string and apply hash in here
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

//VerifyPass compare a password and a hash, if be equal the func return
func VerifyPass(passWithHash, passString string) error {
	return bcrypt.CompareHashAndPassword([]byte(passWithHash), []byte(passString))
}