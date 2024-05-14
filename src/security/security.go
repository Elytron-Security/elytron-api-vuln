package security

import "golang.org/x/crypto/bcrypt"

func Hash(pass string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
}

func CheckPass(stringHash, stringPass string) error {
	return bcrypt.CompareHashAndPassword([]byte(stringHash), []byte(stringPass))
}