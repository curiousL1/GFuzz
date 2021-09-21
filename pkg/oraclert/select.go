package oraclert

import (
	"runtime"
	"sync/atomic"
	"time"
)

// GetSelEfcmCaseIdx will be instrumented to each select in target program.
func GetSelEfcmSwitchCaseIdx(selectID string) int {
	atomic.AddUint32(&getSelEfcmCount, 1)
	idx := efcmStrat.GetCase(selectID)
	if idx != -1 {
		runtime.StoreLastMySwitchChoice(idx)
		return idx
	} else {
		atomic.AddUint32(&notSelEfcmCount, 1)
		runtime.StoreLastMySwitchChoice(-1)
		return -1 // let switch choose the default case
	}
}

func StoreLastMySwitchChoice(choice int) {
	if choice == -1 {
		atomic.AddUint32(&origSelCount, 1)
	}
	runtime.StoreLastMySwitchChoice(choice)
}

func SelEfcmTimeout() <-chan time.Time {
	// if this channel wins, remember to call "runtime.StoreLastMySwitchChoice(-1)", which means we will use the original select
	return time.After(time.Duration(selTimeout) * time.Millisecond)
}
