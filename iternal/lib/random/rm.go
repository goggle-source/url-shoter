package random

import (
	"math/rand"
	"time"
)

func NewRandomString(aliasLength int) string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	arr := []rune("qwertyuiopasdfghjklzxcvbnm" + "QWERTYUIOPASDFGHJKLZXCVBNM" + "1234567890" + "*#$")

	result := make([]rune, aliasLength)

	for i := 0; i < aliasLength; i++ {
		result[i] = arr[rnd.Intn(len(arr))]
	}

	return string(result)
}
