package utils

import (
	"math/rand"
	"time"
)

var RNG *rand.Rand

func init() {
	RNG = rand.New(rand.NewSource(time.Now().UnixNano()))
}
