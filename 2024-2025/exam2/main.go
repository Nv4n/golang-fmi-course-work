package main

import (
	"context"
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func generator(ctx context.Context, power int64) chan int64 {
	outputChan := make(chan int64)
	go func(goCtx context.Context) {
		for i := 1; i <= 1000; i++ {
			randomNum := time.Duration(rand.Int31n(800) + 200)
			select {
			case <-goCtx.Done():
				fmt.Println("Generator goroutine is Done")
				return
			case <-time.After(randomNum * time.Millisecond):
				outputChan <- int64(math.Pow(float64(i), float64(power)))
			}
		}
	}(ctx)
	return outputChan
}

func multiplex(ctx context.Context, channels []chan int64) <-chan int64 {
	var wg sync.WaitGroup
	outputChan := make(chan int64)

	outputFunc := func(id int, inputChan <-chan int64, goCtx context.Context) {
		defer wg.Done()
		for {
			select {
			case <-goCtx.Done():
				fmt.Printf("Ending multiplex goroutine %d\n", id)
				return
			case val, ok := <-inputChan:
				if ok {
					outputChan <- val
				} else {
					return
				}
			}
		}
	}

	for i, inputChan := range channels {
		wg.Add(1)
		go outputFunc(i, inputChan, ctx)
	}
	go func(goCtx context.Context) {
		wg.Wait()
		close(outputChan)

	}(ctx)
	return outputChan
}

func main() {
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	gen1Chan := generator(timeoutCtx, 1)
	gen2Chan := generator(timeoutCtx, 2)
	gen3Chan := generator(timeoutCtx, 3)
	gen4Chan := generator(timeoutCtx, 4)
	gen5Chan := generator(timeoutCtx, 5)

	defer cancel()
	defer close(gen1Chan)
	defer close(gen2Chan)
	defer close(gen3Chan)
	defer close(gen4Chan)
	defer close(gen5Chan)

	channels := []chan int64{gen1Chan, gen2Chan, gen3Chan, gen4Chan, gen5Chan}
	outputChan := multiplex(timeoutCtx, channels)

	for {
		select {
		case <-timeoutCtx.Done():
			cancel()
			time.Sleep(1 * time.Second)
			fmt.Println(fmt.Sprintf("Goroutines: %d", runtime.NumGoroutine()))
			return
		case value := <-outputChan:
			fmt.Printf("New value: %d\n", value)
		}
	}

}
