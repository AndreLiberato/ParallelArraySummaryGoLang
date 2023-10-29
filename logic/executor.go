package logic

import (
	"ParallelArraySummary/model"
	"ParallelArraySummary/util"
	"fmt"
	"math"
	"sync"
	"time"
)

func alocateElements(numberElements uint64) *[]model.Element {
	fmt.Println("Inciando alocação de recursos.")
	startExecution := time.Now()
	elements := make([]model.Element, numberElements)
	endExecution := time.Now()
	fmt.Println("Alocação terminada. Duração(ms)", util.CalculateExecutionTime(startExecution, endExecution))
	util.PrintLine()
	return &elements
}

func generateElementes(elements *[]model.Element, T uint64) *[]*[]model.Element {
	fmt.Println("Preparando dados para execução.")
	startExecution := time.Now()
	sliceElements := util.PrepareElements(elements, T)
	endExecution := time.Now()
	fmt.Println("Preparação terminada. Duração(ms):", util.CalculateExecutionTime(startExecution, endExecution))
	util.PrintLine()
	return sliceElements
}

func runProcess(sliceElements *[]*[]model.Element, T uint64, processWaitGroup *sync.WaitGroup, partialResults chan<- model.Result) {
	fmt.Println("Iniciando processamento dos dados.")
	startExecution := time.Now()
	for i := uint64(0); i < T; i++ {
		processWaitGroup.Add(1)
		go Process((*sliceElements)[i], processWaitGroup, partialResults)
	}
	processWaitGroup.Wait()
	close(partialResults)
	endExecution := time.Now()
	fmt.Println("Processamento de dados terminados. Duração(ms):", util.CalculateExecutionTime(startExecution, endExecution))
	util.PrintLine()
}

func finalProcess(partialResults chan model.Result) *model.Result {
	fmt.Println("Iniciando processamento final dos dados")
	startExecution := time.Now()
	finalResult := getFinalResult(partialResults)
	endExecution := time.Now()
	fmt.Println("Processamento final de dados terminados. Duração(ms):", util.CalculateExecutionTime(startExecution, endExecution))
	util.PrintLine()
	return finalResult
}

func StartProcess(N float64, T uint64) *model.Result {
	var processWaitGroup sync.WaitGroup
	numberElements := uint64(math.Pow(10, N))
	elements := alocateElements(numberElements)

	sliceElements := generateElementes(elements, T)

	partialResults := make(chan model.Result, int(T))

	runProcess(sliceElements, T, &processWaitGroup, partialResults)

	finalResult := finalProcess(partialResults)

	return finalResult
}
