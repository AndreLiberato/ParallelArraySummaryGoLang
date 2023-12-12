package model

type Result struct {
	SumTotal                       float64
	SumTotalByGroup                []float64
	IdsLessThanFiveCounter         uint64
	IdsGreaterOrEqualToFiveCounter uint64
}
