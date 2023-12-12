package logic

import (
	"ParallelArraySummary/model"
	"sync"
)

// SumTotal efetua a soma dos totais de todos os elementos do array
func SumTotal(elements *[]model.Element) float64 {
	countTotal := float64(0)
	for _, element := range *elements {
		countTotal += element.Total
	}
	return countTotal
}

// SumTotalByGroup efetua a soma dos totais dos elementos por grupo
func SumTotalByGroup(elements *[]model.Element) []float64 {
	sumtTotal := []float64{0, 0, 0, 0, 0}
	for _, element := range *elements {
		sumtTotal[element.Group-1] += element.Total
	}
	return sumtTotal
}

func CountTotals(elements *[]model.Element) (uint64, uint64) {
	countTotalLessThanFive := uint64(0)
	countTotalGreaterOrEqualToFive := uint64(0)
	for _, element := range *elements {
		if element.Total >= 5 {
			countTotalGreaterOrEqualToFive++
		} else {
			countTotalLessThanFive++
		}
	}
	return countTotalLessThanFive, countTotalGreaterOrEqualToFive
}

// getFinalResult é a função responsável por processar os resultados parciais dos elementos e gerar o resultado final
func getFinalResult(partialResults chan model.Result) {
	finalResult := model.Result{
		SumTotal:                       0,
		SumTotalByGroup:                []float64{0, 0, 0, 0, 0},
		IdsLessThanFiveCounter:         0,
		IdsGreaterOrEqualToFiveCounter: 0,
	}

	for partialResult := range partialResults {
		finalResult.SumTotal += partialResult.SumTotal
		for group, partialSumByGroup := range partialResult.SumTotalByGroup {
			finalResult.SumTotalByGroup[group] += partialSumByGroup
		}
		finalResult.IdsLessThanFiveCounter += partialResult.IdsLessThanFiveCounter
		finalResult.IdsGreaterOrEqualToFiveCounter += partialResult.IdsGreaterOrEqualToFiveCounter
	}
}

// Process é a função responsável por efetuar sequencialemente as operações sobre os elementos
func Process(elements *[]model.Element, waitGroup *sync.WaitGroup, resultChan chan<- model.Result) {
	defer waitGroup.Done()
	result := model.Result{}

	result.SumTotal = SumTotal(elements)
	result.SumTotalByGroup = SumTotalByGroup(elements)
	result.IdsLessThanFiveCounter, result.IdsGreaterOrEqualToFiveCounter = CountTotals(elements)

	resultChan <- result
}
