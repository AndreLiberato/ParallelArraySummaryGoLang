package util

import "math/rand"

func GenerateGroup() uint8 {
	return uint8(rand.Intn(5) + 1)
}

func GenerateValue() float64 {
	return rand.Float64() * 10
}
