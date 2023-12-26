package helpers

import (
	"math/rand"
	"time"
)

func GenerateID() int64 {
	rand.Seed(time.Now().UnixNano())
	const (
		min = 100000000
		max = 999999999
	)
	return int64(rand.Intn(max-min+1) + min)
}
