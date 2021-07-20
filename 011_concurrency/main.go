package main

import (
	"fmt"
	"runtime"
	"sync"

	"github.com/foroozf001/Golang/011_concurrency/Exercise1"
	"github.com/foroozf001/Golang/011_concurrency/Exercise2"
)

var wg sync.WaitGroup

func main() {
	// Ex1()
	Exercise2.Ex2()
}

// Exercise 1
func Ex1() {
	wg.Add(2)
	go Exercise1.Gor1(GorPrint, &wg)
	go Exercise1.Gor2(GorPrint, &wg)
	NoGor()
	wg.Wait()
}

func NoGor() {
	GorPrint(&wg)
	i := 0
	for i < 100 {
		fmt.Println("NoGor:", i)
		i++
	}
}

func GorPrint(wg *sync.WaitGroup) {
	fmt.Printf("Type:%T, OS:%s, numCPUs:%d,numGoroutines:%d\n", wg, runtime.GOOS, runtime.NumCPU(), runtime.NumGoroutine())
}
