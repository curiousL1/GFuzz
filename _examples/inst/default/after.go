package defaultp

import (
	oraclert "gfuzz/pkg/oraclert"
	aaa "sync"
)

type aa struct{}

func (_ *aa) abcde() {
	println(3)
}

func (a *aa) abcd() {
	oraclert.CurrentGoAddValue(a, nil, 0)
	b := 2
	println(a, b)
}
func Hello() {
	m := aaa.Mutex{}

	c := aaa.NewCond(&m)
	oraclert.StoreOpInfo("Broadcast",

		//asdfasdf
		1)

	c.Broadcast()
	oraclert.StoreOpInfo("Signal", 2)

	c.Signal()
	oraclert.StoreOpInfo("Wait",

		//asdfadfas
		3)

	c.Wait()

	w := aaa.WaitGroup{}
	oraclert.StoreOpInfo("Wait", 4)

	w.Wait()
}
