package util

import (
	"ParallelArraySummaryGoLang/model"
	"math"
)

func LoadElements(N float64) []model.Element {
	var listElements []model.Element
	var numberOfElements uint64 = uint64(math.Pow(10, N))
	var counter uint64 = 0
	for counter < numberOfElements {
		var element model.Element = model.Element{Id: counter + 1, Total: GenerateValue(), Group: GenerateGroup()}
		listElements = append(listElements, element)
		counter++
	}
	return listElements
}
