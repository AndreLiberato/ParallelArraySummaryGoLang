package main

import (
	"ParallelArraySummaryGoLang/logic"
	"ParallelArraySummaryGoLang/util"
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	N, errf := strconv.ParseFloat(args[1], 64)
	T, errui := strconv.ParseUint(args[2], 10, 64)

	if errf != nil {
		fmt.Println("Erro ao converter a string em ponto flutuante:", errf)
	}

	if errui != nil {
		fmt.Println("Erro ao converter a string em ponto inteiro não assinado:", errui)
	}

	result := logic.StartProcess(N, T)

	util.ShowResult(*result)
}
