package logic

import (
	"ParallelArraySummary/model"
	"sync"
)

// SumTotal efetua a soma dos totais de todos os elementos do array
func SumTotal(elements *[]model.Element) float64 {
	var total float64 = 0
	for _, element := range *elements {
		total += element.Total
	}
	return total
}

// SumTotalByGroup efetua a soma dos totais dos elementos por grupo
func SumTotalByGroup(elements *[]model.Element) map[uint8]float64 {
	sumTotalMap := map[uint8]float64{1: 0, 2: 0, 3: 0, 4: 0, 5: 0}
	for _, element := range *elements {
		sumTotalMap[element.Group] += element.Total
	}
	return sumTotalMap
}

// FilterIDByTotalLessThanFive sepera em um aray os ids dos elementos que tem o total são menores do que cinco
func FilterIdByTotalLessThanFive(elements *[]model.Element) []uint64 {
	var idsLessThanFive []uint64
	for _, element := range *elements {
		if element.Total < 5 {
			idsLessThanFive = append(idsLessThanFive, element.Id)
		}
	}
	return idsLessThanFive
}

// FilterIdByTotalGreaterOrEqualToFive separa em um array os ids dos elementos que tem o total maior ou igual que cinco
func FilterIdByTotalGreaterOrEqualToFive(elements *[]model.Element) []uint64 {
	var idsGreaterOrEqualToFive []uint64
	for _, element := range *elements {
		if element.Total >= 5 {
			idsGreaterOrEqualToFive = append(idsGreaterOrEqualToFive, element.Id)
		}
	}
	return idsGreaterOrEqualToFive
}

// getFinalResult é a função responsável por processar os resultados parciais dos elementos e gerar o resultado final
func getFinalResult(partialResults chan model.Result) {
	finalResult := model.Result{}
	finalResult.SumTotalByGroup = map[uint8]float64{1: 0, 2: 0, 3: 0, 4: 0, 5: 0}

	for partialResult := range partialResults {
		finalResult.SumTotal += partialResult.SumTotal
		for group, partialSumByGroup := range partialResult.SumTotalByGroup {
			finalResult.SumTotalByGroup[group] += partialSumByGroup
		}
		finalResult.IdsLessThanFive = append(finalResult.IdsLessThanFive, partialResult.IdsLessThanFive...)
		finalResult.IdsGreaterOrEqualToFive = append(finalResult.IdsGreaterOrEqualToFive, partialResult.IdsGreaterOrEqualToFive...)
	}
}

// Process é a função responsável por efetuar sequencialemente as operações sobre os elementos
func Process(elements *[]model.Element, waitGroup *sync.WaitGroup, resultChan chan<- model.Result) {
	defer waitGroup.Done()
	result := model.Result{}
	result.SumTotal = SumTotal(elements)
	result.SumTotalByGroup = SumTotalByGroup(elements)
	result.IdsLessThanFive = FilterIdByTotalLessThanFive(elements)
	result.IdsGreaterOrEqualToFive = FilterIdByTotalGreaterOrEqualToFive(elements)
	resultChan <- result
}
