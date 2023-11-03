package util

import (
	"ParallelArraySummary/model"
	"fmt"
	"log"
	"os"
	"strings"
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
	fmt.Println("Soma total dos elementos:", result.SumTotal)

	fmt.Println("Soma dos totais dos elementos por grupo:")

	for group, total := range result.SumTotalByGroup {
		fmt.Println("Grupo:", group, "; Total:", total)
	}

	idLess, err := os.Create("id_less_five.csv")
	if err != nil {
		log.Fatalf("An error occurred with file creation: %v", err)
	}
	defer idLess.Close()

	fmt.Println("Escrevendo os id's dos elementos menores que cinco:", len(result.IdsLessThanFive), "elementos.")

	idLess.WriteString(strings.Join(strings.Fields(fmt.Sprint(result.IdsLessThanFive)), "\n"))

	idGreater, err := os.Create("id_greater_five.csv")
	if err != nil {
		log.Fatalf("An error occurred with file creation: %v", err)
	}
	defer idLess.Close()

	fmt.Println("Escrevendo id's dos elementos maiores ou iguais a cinco:", len(result.IdsGreaterOrEqualToFive), "elementos.")

	idGreater.WriteString(strings.Join(strings.Fields(fmt.Sprint(result.IdsGreaterOrEqualToFive)), "\n"))

	fmt.Println("Elementos registrados: ", (len(result.IdsLessThanFive) + len(result.IdsGreaterOrEqualToFive)))
}

func PrintLine() {
	fmt.Println("-------------------------------------------------")
}
