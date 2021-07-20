package Exercise1

import (
	"fmt"
	"sync"
)

func Gor1(f func(wg *sync.WaitGroup), wg *sync.WaitGroup) {
	f(wg)
	i := 0
	for i < 100 {
		fmt.Println("Gor1:", i)
		i++
	}
	wg.Done()
}

func Gor2(f func(wg *sync.WaitGroup), wg *sync.WaitGroup) {
	f(wg)
	i := 0
	for i < 100 {
		fmt.Println("Gor2:", i)
		i++
	}
	wg.Done()
}
