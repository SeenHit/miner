package main

import (
	"fmt"
	"github.com/hacash/core/sys"
	"github.com/hacash/miner/minerworker"
	"github.com/hacash/miner/workerCPU"
	"github.com/hacash/miner/workerGPU"
	"os"
	"os/signal"
	"time"
)

/**

go build -ldflags '-w -s' -o miner_worker_2021_3_22 github.com/hacash/miner/run/minerworker


TEST:

cd ./x16rs/opencl && node pkgclfilego.js && cd ../../ && go build -ldflags '-w -s' -o hacash_miner_worker_2021_03_24_01  miner/run/minerworker/main.go && ./hacash_miner_worker_2021_03_24_01

*/

func main() {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)

	target_ini_file := "minerworker.config.ini"
	// target_ini_file := "/home/shiqiujie/Desktop/Hacash/go/src/github.com/hacash/miner/run/minerworker/test.ini"
	// target_ini_file := ""
	if len(os.Args) >= 2 {
		target_ini_file = os.Args[1]
	}

	target_ini_file = sys.AbsDir(target_ini_file)

	if target_ini_file != "" {
		fmt.Println("Load ini config file: \"" + target_ini_file + "\" at time:" + time.Now().Format("01/02 15:04:05"))
	}

	hinicnf, _ := sys.LoadInicnf(target_ini_file)

	// miner worker
	cnf := minerworker.NewMinerWorkerConfig(hinicnf)
	worker := minerworker.NewMinerWorker(cnf)

	if cnf.GPU_Enable {
		// gpu worker
		gpucnf := workerGPU.NewGpuWorkerConfig(hinicnf)
		gpuworker := workerGPU.NewGpuWorker(gpucnf)
		worker.SetPowWorker(gpuworker)

	} else {
		// cpu worker
		cpucnf := workerCPU.NewCPUWorkerConfig(hinicnf)
		cpuworker := workerCPU.NewCPUWorker(cpucnf)
		worker.SetPowWorker(cpuworker)
	}

	worker.Start()

	s := <-c
	fmt.Println("Got signal:", s)

}
