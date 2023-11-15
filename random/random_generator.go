package random

import "math/rand"

// GenerateGroup gera aleatóriamente um grupo variando no intervalo [1,5]
func GenerateGroup() uint8 {
	return uint8(rand.Intn(5) + 1)
}

// GenerateValue gera aleatóriamente um valor variando no intervalo [0,10[
func GenerateValue() float64 {
	return rand.Float64() * 10
}
