package utils

import (
	"fmt"
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomUser() string {
	return RandomString(6)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}

var accountTypes []string = []string{"debit", "credit"}

func RandomAccountType() string {
	index := RandomInt(0, 1)
	return accountTypes[index]
}
