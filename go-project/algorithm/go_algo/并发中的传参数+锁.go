package go_algo

import (
	"fmt"
	"sync"
	"time"
)

func main1() {
	s := []int{1, 2, 3}

	for _, v := range s {
		//a := v
		go func(v int) {

			println(v)

		}(v)
	}

	//select {}
	time.Sleep(1 * time.Second)
}
func main2() {
	var mu sync.Mutex
	counter := 0

	for i := 0; i < 1000; i++ {

		go func() {
			mu.Lock()
			counter++
			mu.Unlock()

		}()
	}
	time.Sleep(time.Second)
	fmt.Println("Finally Counter:", counter)
	//select {}
}
