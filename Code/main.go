package main

import (
	"fmt"
	"sync"
)

func raceCondition() int {
	data := 0
	var dataAccess sync.Mutex
	go func() {
		dataAccess.Lock()
		data++
		dataAccess.Unlock()
	}()
	dataAccess.Lock()
	if data == 0 {
		print("")
		return data
	}
	dataAccess.Unlock()
	return -1
}

func print_shares(values []int, outputs []int) {
	out_len := len(outputs)
	for _, val := range values {
		counter := 0
		for _, out := range outputs {
			if out == val {
				counter++
			}
		}
		fmt.Printf("Counts of %d: %d\n", val, counter)
		fmt.Printf("Share of %d: %f Percent\n", val, (float64(counter)/float64(out_len))*100)
	}
}

func showRaceCondition() {
	var outputs []int
	for i := 0; i < 100000; i++ {
		outputs = append(outputs, raceCondition())
	}
	print_shares([]int{-1, 0, 1}, outputs)
}

func simple_deadlock() {
	x := 0
	y := 0
	var xMutex sync.Mutex
	var yMutex sync.Mutex
	go func() {
		xMutex.Lock()
		yMutex.Lock()
		x, y = doWork(x, y)
		yMutex.Unlock()
		xMutex.Unlock()
	}()
	yMutex.Lock()
	xMutex.Lock()
	x, y = doOtherWork(x, y)
	xMutex.Unlock()
	yMutex.Unlock()
}

func doWork(a int, b int) (int, int) {
	if a == 1 && b == 1 {
		a++
		b++
	}
	return a, b
}

func doOtherWork(a int, b int) (int, int) {
	if a == 0 && b == 0 {
		b++
		a++
	}
	return a, b
}

func communication_deadlock() {
	ch := make(chan int)

	go func() {
		data := <-ch
		fmt.Println("Received:", data)
	}()

	data := <-ch
	fmt.Println("Received:", data)
}

func deadlock3Threads() {
	a := 0
	b := 0
	c := 0
	lockA := &sync.Mutex{}
	lockB := &sync.Mutex{}
	lockC := &sync.Mutex{}

	go func() {
		lockA.Lock()
		lockB.Lock()
		if a == 0 && b == 0 {
			a++
			b++
		}
		lockB.Unlock()
		lockA.Unlock()
	}()

	go func() {
		lockB.Lock()
		lockC.Lock()
		if b == 0 && c == 0 {
			b++
			c++
		}
		lockC.Unlock()
		lockB.Unlock()
	}()

	lockC.Lock()
	lockA.Lock()
	if c == 0 && a == 0 {
		c++
		a++
	}
	lockA.Unlock()
	lockC.Unlock()
}

func main() {
	for i := 0; i < 10000000; i++ {
		deadlock3Threads()
	}
}
