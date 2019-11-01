package rand

import (
	"math/rand"
	"time"
)

func RandomInt(min int, max int) int {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	return random.Intn(max+1-min) + min
}
