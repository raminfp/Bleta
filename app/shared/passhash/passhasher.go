package passhash

import (
	"golang.org/x/crypto/bcrypt"
)

// source lib : https://github.com/golang/crypto/blob/master/bcrypt/bcrypt.go

// HashString returns a hashed string and an error
func HashString(password string) (string, error) {
	key, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(key), nil
}

// GenerateFromPassword returns the bcrypt hash of the password at the given
func HashBytes(password []byte) ([]byte, error) {
	key, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return key, nil
}

// CompareHashAndPassword compares a bcrypt hashed password with its possible
// plaintext equivalent. Returns nil on success, or an error on failure.

func MatchString(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err == nil {
		return true
	}

	return false
}

// MatchBytes returns true if the hash matches the password
func MatchBytes(hash, password []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, []byte(password))
	if err == nil {
		return true
	}
	return false
}