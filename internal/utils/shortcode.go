package utils

import "crypto/rand"

const charSet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateRandomCode(length int) (string, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)

	if err != nil {
		return "", err
	}
	for i := range b {
		b[i] = charSet[int(b[i])%len(charSet)]
	}

	return string(b), nil
}
