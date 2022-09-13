package pattern

import (
	"fmt"
	"sync"
)

func FaninMain() {
	input1 := GenerateWork([]int{0, 2, 4, 6, 8, 10, 12})
	input2 := GenerateWork([]int{1, 3, 5, 7, 9, 11, 13})

	out := fanIn(input1, input2)
	for value := range out {
		fmt.Println("Value: ", value)
	}
}

func fanIn(inputs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	wg.Add(len(inputs))

	for _, in := range inputs {
		go func(ch <-chan int) {
			for {
				value, ok := <-ch
				if !ok {
					wg.Done()
					break
				}
				out <- value
			}
		}(in)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func GenerateWork(work []int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)

		for _, w := range work {
			ch <- w
		}
	}()
	return ch
}
