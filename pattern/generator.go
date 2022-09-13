package pattern

import "fmt"

func GeneratorMain() {
	size := 20
	ch := generatorPattern(size)

	for i := 0; i < size; i++ {
		value := <-ch
		fmt.Println(value)
	}
}

func generatorPattern(size int) <-chan int {
	ch := make(chan int, size)

	go func() {
		for i := 0; i < size; i++ {
			ch <- i
		}
	}()

	return ch
}
