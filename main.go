package main

import (
	"ParallelArraySummary/handler"
	"ParallelArraySummary/logic"
	"fmt"
	"os"
	"runtime/pprof"
	"strconv"
	"time"
)

func main() {
	fcpu, err := os.Create("cpu.prof")
	handler.Chk("Erro ao criar o arquivo de prof", err)
	defer fcpu.Close()
	pprof.StartCPUProfile(fcpu)
	defer pprof.StopCPUProfile()

	startTime := time.Now() // Marcador de tempo de início do programa

	args := os.Args
	N, errf := strconv.ParseFloat(args[1], 64) // N expoente
	handler.Chk("Erro ao converter a string em ponto flutuante:", errf)

	T, errui := strconv.ParseUint(args[2], 10, 64) // Número de threads
	handler.Chk("Erro ao converter a string em inteiro:", errui)

	logic.StartProcess(N, T) //Inicio do processo

	endTime := time.Now() // Marcador de tempo de fim do programa

	programDurationTime := endTime.Sub(startTime).Milliseconds()

	fmt.Println("Duração da execução total do programa:", programDurationTime)

	fmem, err := os.Create("mem.prof")
	handler.Chk("Erro ao criar o arquivo de prof", err)
	defer fmem.Close()
	pprof.WriteHeapProfile(fmem)
}
