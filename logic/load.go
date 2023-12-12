package logic

import (
	"ParallelArraySummary/model"
	"ParallelArraySummary/random"
)

// loadElements é a função responsável por gerar o array de elementos aleatórios para cada thread
func loadElements(subListSize uint64, idCounter uint64) *[]model.Element {
	subList := make([]model.Element, subListSize)
	for j := uint64(0); j < subListSize; j++ {
		idCounter++
		subList[j].Id = idCounter
		subList[j].Total = random.GenerateValue()
		subList[j].Group = random.GenerateGroup()
	}
	return &subList
}

// PrepareElements é a função responsável por criar os elementos por thread de maneira equilibrada
func PrepareElementsLinear(numberElements uint64, T uint64) *[]*[]model.Element {
	chunkSize := numberElements / T              // Quantidade de elementos por thread
	extraElements := numberElements % T          // Quantidade de elementos restantes
	sliceElements := make([]*[]model.Element, T) // Array de ponteiros para arrays de Element. Armazena um ponteiro para array por thread.
	idCounter := uint64(0)
	for i := uint64(0); i < T; i++ {
		subListSize := chunkSize
		if i < extraElements {
			subListSize++
		}
		sliceElements[i] = loadElements(subListSize, idCounter)
		idCounter += subListSize
	}
	return &sliceElements
}
