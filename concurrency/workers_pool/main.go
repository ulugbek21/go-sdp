package main

import (
	"fmt"
	"sync"
)

func main() {
	bufferSize := 100
	var dispatcher Dispatcher = NewDispatcher(bufferSize)

	workers := 3
	for i := 0; i < workers; i++ {
		var w WorkerLauncher = &PreffixSuffixWorker{
			prefixS: fmt.Sprintf("WorkerID: %d -> ", i),
			suffixS: " World",
			id:      i,
		}
		dispatcher.LaunchWorker(w)
	}

	requests := 10

	var wg sync.WaitGroup
	wg.Add(requests)

	for i := 0; i < requests; i++ {
		req := NewStringRequest(fmt.Sprintf("(Msg_id: %d) -> Hello", i), i, &wg)
		dispatcher.MakeRequest(req)
	}

	dispatcher.Stop()

	wg.Wait()
}
