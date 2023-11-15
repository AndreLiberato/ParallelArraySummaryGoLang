package logic

import (
	"ParallelArraySummary/model"
	"ParallelArraySummary/random"
	"sync"
)

// loadElements é a função responsável por gerar o array de elementos aleatórios para cada thread
func loadElements(sliceElements **[]model.Element, listSize uint64, idCounter *uint64, waitGroup *sync.WaitGroup, idMutex *sync.Mutex) {
	defer waitGroup.Done()                  // Pornto de sincronia
	list := make([]model.Element, listSize) // Alocação do array de elementos
	for i := uint64(0); i < listSize; i++ {
		list[i].Id = getNextID(idCounter, idMutex)
		list[i].Total = random.GenerateValue()
		list[i].Group = random.GenerateGroup()
	}
	*sliceElements = &list
}

// getNextID é a função responsável por gerar o próximo id sequencial
func getNextID(idCounter *uint64, idMutex *sync.Mutex) uint64 {
	idMutex.Lock()   // Bloqueia o recurso
	*idCounter += 1  // Adiciona + 1
	idMutex.Unlock() // Desbloqueia o recurso
	return *idCounter
}

// PrepareElements é a função responsável por criar os elementos por thread de maneira equilibrada
func PrepareElements(numberElements uint64, T uint64) *[]*[]model.Element {
	chunkSize := numberElements / T // Quantidade de elementos por thread

	extraElements := numberElements % T // Quantidade de elementos restantes

	var loadElementsWaitGroup sync.WaitGroup // Variaǘel de sincronia final do carregamento

	var idMutex sync.Mutex // Mutex para geração sequencial de ids

	var idCounter uint64 // Contador de ids

	var sliceElements []*[]model.Element = make([]*[]model.Element, T) // Array de ponteiros para arrays de Element. Armazena um ponteiro para array por thread.

	for i := uint64(0); i < T; i++ {
		loadElementsWaitGroup.Add(1)

		subListSize := chunkSize
		if i < extraElements {
			subListSize++
		}

		// Inicia as goroutines
		go loadElements(&sliceElements[i], subListSize, &idCounter, &loadElementsWaitGroup, &idMutex)
	}

	loadElementsWaitGroup.Wait() // Espera o término da geração de elementos

	return &sliceElements
}
