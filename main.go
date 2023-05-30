package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/shirou/gopsutil/process"
)

func main() {
	processName := "CalculatorApp.exe"
	processes, err := process.Processes()
	if err != nil {
		fmt.Println("Erro ao listar processos:", err)
		os.Exit(1)
	}

	found := false
	for _, p := range processes {
		executable, err := p.Exe()
		if err != nil {
			continue
		}
		if strings.Contains(strings.ToLower(executable), strings.ToLower(processName)) {
			found = true
			fmt.Println("Processo encontrado:")
			fmt.Println("PID:", p.Pid)
			fmt.Println("Nome:", executable)

			// Obtendo informações adicionais do processo
			memInfo, err := p.MemoryInfo()
			if err == nil {
				fmt.Println("Uso de Memória:", memInfo.RSS)
			}

			cpuPercent, err := p.CPUPercent()
			if err == nil {
				fmt.Println("Uso de CPU:", cpuPercent)
			}

			createTime, err := p.CreateTime()
			if err == nil {
				fmt.Println("Tempo de criação:", createTime)
			}

			threads, err := p.NumThreads()
			if err == nil {
				fmt.Println("Número de Threads:", threads)
			}
			time.Sleep(5 * time.Second)

			fmt.Println("Encerrando processo...")
			err = p.Kill()
			if err != nil {
				fmt.Println("Erro ao encerrar processo:", err)
			}
			break
		}
	}

	if !found {
		fmt.Println("Processo não encontrado")
	}
}