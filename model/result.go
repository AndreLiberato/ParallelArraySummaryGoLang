package model

type Result struct {
	SumTotal                float64
	SumTotalByGroup         map[uint8]float64
	IdsLessThanFive         []uint64
	IdsGreaterOrEqualToFive []uint64
}
