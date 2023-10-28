package util

import (
	"ParallelArraySummaryGoLang/model"
)

func LoadElements(numberElements uint64) *[]model.Element {
	var listElements []model.Element
	var counter uint64 = 0
	for counter < numberElements {
		var element model.Element = model.Element{Id: counter + 1, Total: GenerateValue(), Group: GenerateGroup()}
		listElements = append(listElements, element)
		counter++
	}
	return &listElements
}

func SliceElements(elements *[]model.Element, T uint64) *[][]model.Element {
	numberElements := uint64(len(*elements))
	chunkSize := numberElements / T
	extraElements := numberElements % T
	var startIndex uint64 = 0
	var sliceElements [][]model.Element
	for i := uint64(0); i < T; i++ {
		subListSize := chunkSize
		if i < extraElements {
			subListSize++
		}
		endIndex := startIndex + subListSize
		sliceElements = append(sliceElements, (*elements)[startIndex:endIndex])
		startIndex = endIndex
	}
	return &sliceElements
}
