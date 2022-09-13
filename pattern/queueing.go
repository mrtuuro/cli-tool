package pattern

import (
	"fmt"
	"sync"
	"time"
)

const limit = 8
const work = 450

func QueueingMain() {
	var wg sync.WaitGroup
	fmt.Println("Queue limit: ", limit)
	queue := make(chan struct{}, limit)

	wg.Add(work)

	for w := 1; w <= work; w++ {
		process(w, queue, &wg)
	}

	wg.Wait()

	close(queue)
	fmt.Println("Work complete")
}

func process(work int, queue chan struct{}, wg *sync.WaitGroup) {
	queue <- struct{}{}

	go func() {
		defer wg.Done()

		time.Sleep(100 * time.Millisecond)
		fmt.Println("Processed: ", work)

		<-queue
	}()
}
