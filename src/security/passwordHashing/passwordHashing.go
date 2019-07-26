package passwordHashing

import (
	"golang.org/x/crypto/bcrypt"
	"errors"
	"log"
)

func hashPassword(password string) (string, error){
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}

func ComparePasswords(hashedPassword string, plainPwd []byte) bool {
	byteHash := []byte(hashedPassword)

	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func HashAndCheckPassword(password string) (string, error){
	if hash, hashError := hashPassword(password); hashError == nil{
		if err := checkPasswordHash(password, hash); err == nil{
			return hash, err
		}
		return "", errors.New("Password hash check failed")
	}
	return "", errors.New("Hash failed")
}