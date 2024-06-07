package utils

import (
	"math/rand"
	"time"
)

func Choice(s []string) string {
	rand.Seed(time.Now().UnixNano())
	return s[rand.Intn(len(s))]
}
