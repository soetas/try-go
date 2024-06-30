package rand

import (
	"fmt"
	"math/rand"
	"time"
)

func Color() string {
	rand.Seed(time.Now().UnixNano())

	r := rand.Intn(256)
	g := rand.Intn(256)
	b := rand.Intn(256)

	return fmt.Sprintf("0x%x%x%x", r, g, b)
}
