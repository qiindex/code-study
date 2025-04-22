package three_print_abc

import (
	"sync"
)

var count = 5

func Main2() {

	wg := sync.WaitGroup{}
	chanA := make(chan int, 1)
	chanB := make(chan int, 1)
	chanC := make(chan int, 1)

	chanA <- 1
	wg.Add(3)

	go printA(&wg, chanA, chanB)
	go printB(&wg, chanB, chanC)
	go printC(&wg, chanC, chanA)
	wg.Wait()
}

func printA(wg *sync.WaitGroup, chanA chan int, chanB chan int) {
	defer wg.Done()

	for i := 0; i < count; i++ {
		_ = <-chanA
		println("A")
		chanB <- 1
	}
}

func printB(wg *sync.WaitGroup, chanB chan int, chanC chan int) {
	defer wg.Done()

	for i := 0; i < count; i++ {
		_ = <-chanB
		println("B")
		chanC <- 2
	}
}

func printC(wg *sync.WaitGroup, chanC chan int, chanA chan int) {
	defer wg.Done()

	for i := 0; i < count; i++ {
		<-chanC
		println("C")
		chanA <- 3
	}
}
