package security

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func GeneratePassKeyHash(username string) (string, error) {
	// create a random 64 bytes (512 bits) secret
	secret := make([]byte, 64)
	_, err := rand.Read(secret)
	if err != nil {
		fmt.Println("error generating a random secret:", err)
		return "", err
	}

	data := []byte(username + "salt")

	// create a new HMAC by defining the hash type and the key
	hmac := hmac.New(sha256.New, secret)

	// compute the HMAC
	hmac.Write([]byte(data))
	dataHmac := hmac.Sum(nil)

	hmacHex := hex.EncodeToString(dataHmac)
	return hmacHex, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
