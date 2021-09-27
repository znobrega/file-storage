package helpers

import "golang.org/x/crypto/bcrypt"

func CheckPassowordHash(password, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil, err
}
