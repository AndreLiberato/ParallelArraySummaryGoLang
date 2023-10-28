package logic

import (
	"ParallelArraySummaryGoLang/model"
	"ParallelArraySummaryGoLang/util"
	"fmt"
	"math"
	"sync"
	"time"
)

func StartProcess(N float64, T uint64) *model.Result {
	var waitGroup sync.WaitGroup
	partialResults := make(chan model.Result, int(T))

	elements := util.LoadElements(uint64(math.Pow(10, N)))

	sliceElements := util.SliceElements(elements, T)

	startTime := time.Now()

	for i := 0; i < len(*sliceElements); i++ {
		waitGroup.Add(1)
		go Process(&(*sliceElements)[i], &waitGroup, partialResults)
	}

	waitGroup.Wait()

	close(partialResults)

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
	endTime := time.Now()

	totalTime := endTime.Sub(startTime)

	timeInMilliseconds := totalTime.Milliseconds()

	fmt.Println("Duração total do processamento:", timeInMilliseconds, "ms")

	return &finalResult
}
