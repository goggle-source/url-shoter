package random

import (
	"math/rand"
	"time"
)

func NewRandomString(size int) string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	arr := []rune("qwertyuiopasdfghjklzxcvbnm" + "QWERTYUIOPASDFGHJKLZXCVBNM" + "1234567890" + "*#$")

	result := make([]rune, size)

	for i := range result {
		result[i] = arr[rnd.Intn(len(arr))]
	}

	return string(result)
}
