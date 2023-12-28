package concurrency

import (
	"sync"
)

func OrWithWG(channels []<-chan any) <-chan any {
	sink := make(chan any, 1)
	var wg sync.WaitGroup
	wg.Add(len(channels))

	for _, channel := range channels {
		go func(ch <-chan any) {
			defer wg.Done()
			for v := range ch {
				sink <- v
			}
		}(channel)

	}

	go func() {
		wg.Wait()
		close(sink)
	}()

	return sink
}

func OrWithChannel(channels []<-chan any) <-chan any {
	sink := make(chan any)
	sem := make(chan struct{}, len(channels))

	for _, channel := range channels {
		sem <- struct{}{}
		go func(ch <-chan any) {
			for v := range ch {
				sink <- v
			}
			<-sem
		}(channel)

	}

	go func() {
		for i := 0; i < len(channels); i++ {
			sem <- struct{}{}
		}
		close(sink)
	}()

	return sink
}
