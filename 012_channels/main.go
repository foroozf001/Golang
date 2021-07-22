package main

import (
	"fmt"
)

func main() {
	// Exercise1()
	// Exercise2a()
	// Exercise2b()
	// Exercise3()
	// Exercise4()
	// Exercise5()
	// Exercise6()
	Exercise7()
}

func Exercise1() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}

func Exercise2a() {
	cs := make(chan int)

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}

func Exercise2b() {
	cr := make(chan int)

	go func() {
		cr <- 42
	}()
	fmt.Println(<-cr)

	fmt.Printf("------\n")
	fmt.Printf("cr\t%T\n", cr)
}

func Exercise3() {
	c := gen()

	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

func gen() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 20; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}

func Exercise4() {
	q := make(chan int) //bi-directional
	c := gen4(q)        //receive-only

	receive4(c, q)

	fmt.Println("about to exit")
}

func receive4(c <-chan int, q chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-q:
			return
		}
	}
}

func gen4(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		q <- 1
		close(c)
	}()

	return c
}

func Exercise5() {
	c := make(chan int)

	go func() {
		c <- 1
		c <- 2
	}()

	v, ok := <-c
	fmt.Println(v, ok)

	close(c)

	v, ok = <-c
	fmt.Println(v, ok)
}

func Exercise6() {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

func Exercise7() {
	c := make(chan int)

	for j := 0; j < 10; j++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
		}()
	}

	for k := 0; k < 100; k++ {
		fmt.Println(k, <-c)
	}

	fmt.Println("about to exit")
}
