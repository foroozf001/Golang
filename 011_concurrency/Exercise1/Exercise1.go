package Exercise1

import (
	"sync"
)

func Gor1(f func(wg *sync.WaitGroup), wg *sync.WaitGroup) {
	f(wg)
	wg.Done()
}

func Gor2(f func(wg *sync.WaitGroup), wg *sync.WaitGroup) {
	f(wg)
	wg.Done()
}
