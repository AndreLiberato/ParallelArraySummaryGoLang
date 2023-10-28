package logic

import (
	"ParallelArraySummaryGoLang/model"
	"sync"
)

func SumTotal(elements *[]model.Element) float64 {
	var total float64 = 0
	for _, element := range *elements {
		total += element.Total
	}
	return total
}

func SumTotalByGroup(elements *[]model.Element) map[uint8]float64 {
	sumTotalMap := map[uint8]float64{1: 0, 2: 0, 3: 0, 4: 0, 5: 0}

	for _, element := range *elements {
		sumTotalMap[element.Group] += element.Total
	}

	return sumTotalMap
}

func FilterIDByTotalLessThanFive(elements *[]model.Element) []uint64 {
	var idsLessThanFive []uint64

	for _, element := range *elements {
		if element.Total < 5 {
			idsLessThanFive = append(idsLessThanFive, element.Id)
		}
	}

	return idsLessThanFive
}

func FilterIdByTotalGreaterOrEqualToFive(elements *[]model.Element) []uint64 {
	var idsGreaterOrEqualToFive []uint64

	for _, element := range *elements {
		if element.Total >= 5 {
			idsGreaterOrEqualToFive = append(idsGreaterOrEqualToFive, element.Id)
		}
	}

	return idsGreaterOrEqualToFive
}

func Process(elements *[]model.Element, waitGroup *sync.WaitGroup, resultChan chan<- model.Result) {
	defer waitGroup.Done()
	result := model.Result{}
	result.SumTotal = SumTotal(elements)
	result.SumTotalByGroup = SumTotalByGroup(elements)
	result.IdsLessThanFive = FilterIDByTotalLessThanFive(elements)
	result.IdsGreaterOrEqualToFive = FilterIdByTotalGreaterOrEqualToFive(elements)
	resultChan <- result
}
