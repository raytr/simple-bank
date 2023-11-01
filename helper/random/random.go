package random

import "math/rand"

const letterAndNumberBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateRandomString(length int) string {
	var res []byte
	for i := 0; i < length; i++ {
		res = append(res, letterAndNumberBytes[rand.Intn(len(letterAndNumberBytes))])
	}
	return string(res)
}
