package fuzzer

import (
	"context"
	"fmt"
	gLog "gfuzz/pkg/fuzzer/log"
	"log"
	"strconv"
	"sync"
)

// startWorkers starts parallel workers working on inputCh from fuzzer context.
func startWorkers(parallel int, worker func(context.Context)) {
	var wg sync.WaitGroup
	for i := 1; i <= parallel; i++ {
		wg.Add(1)
		ctx := newWorkerContext(strconv.Itoa(i))

		// Start worker
		go func(workerID int) {
			defer wg.Done()
			worker(ctx)
		}(i)
	}

	wg.Wait()
}

const CTX_KEY_WORKER_ID = "WORKER_ID"

func newWorkerContext(workerID string) context.Context {
	return context.WithValue(context.Background(), CTX_KEY_WORKER_ID, workerID)
}

func getWorkerID(context context.Context) string {
	return context.Value(CTX_KEY_WORKER_ID).(string)
}

func getWorkerLogger(context context.Context) *log.Logger {
	workerID := getWorkerID(context)
	return gLog.NewLogger(fmt.Sprintf("[worker %s] ", workerID))
}
