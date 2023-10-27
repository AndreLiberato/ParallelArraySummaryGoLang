package model

type Result struct {
	AllTotal                float64
	SumTotalMap             map[uint8]float64
	IdsLessThanFive         []uint64
	IdsGreaterOrEqualToFive []uint64
}
