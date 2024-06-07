package random

import (
	"math/rand"
	"time"
)

func Choice(slice []interface{}) interface{} {
	rand.Seed(time.Now().UnixNano())
	return slice[rand.Intn(len(slice))]
}
