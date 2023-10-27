package util

import (
	"ParallelArraySummaryGoLang/model"
	"fmt"
)

func ToString(element *model.Element) {
	fmt.Println("Elemento [ id = ", element.Id, ", total = ", element.Total, ", grupo = ", element.Group, " ]")
}

func ArrayToString(elements *[]model.Element) {
	for _, element := range *elements {
		ToString(&element)
	}
}

func ShowResult(result model.Result) {
	fmt.Println("Soma total dos elementos:", result.AllTotal)

	fmt.Println("Soma dos totais dos elementos por grupo:")

	for group, total := range result.SumTotalMap {
		fmt.Println("Grupo:", group, "; Total:", total)
	}

	// fmt.Println("Id's dos elementos menores que cinco:")

	// fmt.Println(strings.Join(strings.Fields(fmt.Sprint(result.IdsLessThanFive)), ", "))

	// fmt.Println("Id's dos elementos maiores ou iguais a cinco:")

	// fmt.Println(strings.Join(strings.Fields(fmt.Sprint(result.IdsGreaterOrEqualToFive)), ", "))
}
