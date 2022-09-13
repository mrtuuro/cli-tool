package pattern

import (
	"fmt"
	"math"
)

func PipelineMain() {
	in := GenerateWork([]int{1, 298, 7, 43, 2, 74, 53, 7, 8, 86, 634, 2111, 412, 4, 512})

	// "in" kanalındaki inputları alıp önce filtreleme stage'ine sonra küp alma stage'ine ve en son olarak
	// yarıya indirme stage'ine sokup çıkan değerleri bastırıyoruz.
	out := filter(in)
	out = third(out)
	out = half(out)

	for value := range out {
		fmt.Println(value)
	}
}

func filter(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for i := range in {
			if i%2 == 0 {
				out <- i
			}

		}
	}()
	return out
}

func third(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for i := range in {
			value := math.Pow(float64(i), 3)
			out <- int(value)
		}
	}()
	return out
}

func half(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for i := range in {
			value := float64(i) / 2
			out <- int(value)
		}

	}()
	return out
}
