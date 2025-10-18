package random_sleep

import (
	"math/rand"
	"time"
)

func RandomSleep(min, max time.Duration) {
	duration := time.Duration(rand.Int63n(int64(max-min+1)) + int64(min))
	time.Sleep(duration)
}
