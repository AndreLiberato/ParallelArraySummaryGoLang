package util

import (
	"ParallelArraySummary/model"
	"sync"
)

func loadElements(sliceElements **[]model.Element, listSize uint64, idCounter *uint64, waitGroup *sync.WaitGroup, idMutex *sync.Mutex) {
	defer waitGroup.Done()
	list := make([]model.Element, listSize)
	for i := uint64(0); i < listSize; i++ {
		list[i].Id = getNextID(idCounter, idMutex)
		list[i].Total = GenerateValue()
		list[i].Group = GenerateGroup()
	}
	*sliceElements = &list
}

func getNextID(idCounter *uint64, idMutex *sync.Mutex) uint64 {
	idMutex.Lock()
	defer idMutex.Unlock()
	*idCounter += 1
	return *idCounter
}

func PrepareElements(elements *[]model.Element, T uint64) *[]*[]model.Element {
	numberElements := uint64(len(*elements))
	chunkSize := numberElements / T
	extraElements := numberElements % T
	var loadElementsWaitGroup sync.WaitGroup
	var idMutex sync.Mutex
	var idCounter uint64
	var sliceElements []*[]model.Element = make([]*[]model.Element, T)

	for i := uint64(0); i < T; i++ {
		loadElementsWaitGroup.Add(1)

		subListSize := chunkSize
		if i < extraElements {
			subListSize++
		}

		go loadElements(&sliceElements[i], subListSize, &idCounter, &loadElementsWaitGroup, &idMutex)
	}

	loadElementsWaitGroup.Wait()
	return &sliceElements
}
