package main

import (
	"fmt"
	"runtime"
	"sync"

	"github.com/foroozf001/Golang/011_concurrency/Exercise1"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	NoGor()
	go Exercise1.Gor1(GorPrint, &wg)
	go Exercise1.Gor2(GorPrint, &wg)
	wg.Wait()
}

func NoGor() {
	GorPrint(&wg)
}

func GorPrint(wg *sync.WaitGroup) {
	fmt.Printf("Type:%T, OS:%s, numCPUs:%d,numGoroutines:%d\n", wg, runtime.GOOS, runtime.NumCPU(), runtime.NumGoroutine())
}
