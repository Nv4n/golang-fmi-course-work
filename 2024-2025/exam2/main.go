package main

import (
	"context"
	"fmt"
	"log"
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
				log.Println("Gen1 Done")
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

	output := func(id int, inputChan <-chan int64) {
		defer wg.Done()

		for value := range inputChan {
			select {
			case <-ctx.Done():
				log.Println("Ending multiplex goroutine")
			case outputChan <- value:

			}

		}
	}

	wg.Add(len(channels))
	for i, inputChan := range channels {
		go output(i, inputChan)
	}
	go func() {
		defer close(outputChan)
		wg.Wait()
	}()
	return outputChan
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	gen1Chan := generator(ctx, 1)
	gen2Chan := generator(ctx, 2)
	gen3Chan := generator(ctx, 3)
	gen4Chan := generator(ctx, 4)
	gen5Chan := generator(ctx, 5)

	defer cancel()
	defer close(gen1Chan)
	defer close(gen2Chan)
	defer close(gen3Chan)
	defer close(gen4Chan)
	defer close(gen5Chan)
	defer log.Println(fmt.Sprintf("Goroutines: %d", runtime.NumGoroutine()))

	channels := []chan int64{gen1Chan, gen2Chan, gen3Chan, gen4Chan, gen5Chan}
	outputChan := multiplex(ctx, channels)

	timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	for {
		select {
		case <-timeoutCtx.Done():
			cancel()
			return
		case value := <-outputChan:
			log.Printf("New value: %d", value)
		}
	}
}
