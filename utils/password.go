package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(s *string) error {
	if s == nil {
		return fmt.Errorf("param was empty.")
	}
	sBytes := []byte(*s)
	hashBytes, err := bcrypt.GenerateFromPassword(sBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	*s = string(hashBytes[:])
	return nil
}

func CheckPassword(existinghash, incomingPass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(existinghash), []byte(incomingPass)) == nil
}
