package random

import (
	"math/rand"
	"time"
)

func GetRandomInt(min int, max int) int {
	defineSeed()
	return rand.Intn(max-min) + min
}

func GetRandomFloat(min float32, max float32) float32 {
	defineSeed()
	return min + rand.Float32()*(max-min)
}

//https://stackoverflow.com/questions/39529364/go-rand-intn-same-number-value
func defineSeed() {
	rand.Seed(time.Now().UnixNano())
}
