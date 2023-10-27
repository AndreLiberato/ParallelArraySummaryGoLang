package main

import (
	"ParallelArraySummaryGoLang/logic"
	"ParallelArraySummaryGoLang/util"
)

func main() {
	var N float64 = 10
	// var T int = 10

	listElements := util.LoadElements(N)

	result := logic.Process(&listElements)

	util.ShowResult(result)
}
