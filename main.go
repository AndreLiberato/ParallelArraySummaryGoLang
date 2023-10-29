package main

import (
	"ParallelArraySummary/logic"
	"ParallelArraySummary/util"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
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

	f, err := os.Create("execution_info.prof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	result := logic.StartProcess(N, T)

	util.ShowResult(*result)
}
