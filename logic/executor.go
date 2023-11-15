package logic

import (
	"ParallelArraySummary/model"
	"fmt"
	"math"
	"sync"
	"time"
)

// runProcess é a função responsável por iniciar as goroutines do processamento dos elementos
func runProcess(sliceElements *[]*[]model.Element, T uint64, processWaitGroup *sync.WaitGroup, partialResults chan<- model.Result) {
	defer close(partialResults)
	for i := uint64(0); i < T; i++ {
		processWaitGroup.Add(1)
		go Process((*sliceElements)[i], processWaitGroup, partialResults)
	}
	processWaitGroup.Wait()
}

// StartProcess é a função responsável pelo fluxo de processamento principal do programa
func StartProcess(N float64, T uint64) {
	var processWaitGroup sync.WaitGroup // Variável de sincronização final

	numberElements := uint64(math.Pow(10, N)) // Número de elementos

	sliceElements := PrepareElements(numberElements, T) // Gera os elementos por thread

	partialResults := make(chan model.Result, T) // Canal de resultados parciais

	startTime := time.Now() // Marcador de tempo de inicio do processamento dos elementos

	runProcess(sliceElements, T, &processWaitGroup, partialResults)

	getFinalResult(partialResults) // Resolve todos os resultados parciais em só um resultado

	endTime := time.Now() // Marcador de tempo de final do processamento dos elementos

	totalTime := endTime.Sub(startTime).Milliseconds() // Tempo total do processamento dos elementos

	fmt.Println("Tempo de processamento total dos elementos:", totalTime)
}
