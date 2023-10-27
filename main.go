package main

import (
	"fmt"
	"strconv"
)

// Elements
type Element struct {
	id    int
	total float32
	group int8
}

func toString(element Element) string {
	var sId string
	var sTotal string
	var sGroup string

	sId = strconv.Itoa(element.id)
	sTotal = strconv.FormatFloat(float64(element.total), 'e', 2, 64)
	sGroup = strconv.Itoa(int(element.group))

	return "Elemento [ id = " + sId + ", total = " + sTotal + ", grupo = " + sGroup + " ]"
}

//ParallelArraySummary.java

var elements []Element

//IOperation

func sumTotal(partElements []Element) float32 {
	var allTotal float32 = 0
	for _, indTotal := range partElements {
		allTotal += indTotal.total
	}
	return allTotal
}

func sumTotalByGroup(partElements []Element) map[int8]float32 {
	sumTotalMap := make(map[int8]float32, 5)
	var null float32
	for _, element := range partElements {
		var sumTotal float32
		if sumTotalMap[element.group] == null {
			sumTotal = 0.0
		} else {
			sumTotal = sumTotalMap[element.group]
		}
		var total = sumTotal + element.total
		sumTotalMap[element.group] = total
	}
	return sumTotalMap
}

func main() {
	element := new(Element)
	element.id = 12
	fmt.Println(toString(*element))
}
