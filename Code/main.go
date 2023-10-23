package main

import (
	"fmt"
)

func race_condition() int {
	data := 0
	go func() {
		data++
	}()
	if data == 0 {
		print("")
		return data
	}
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

func main() {
	var outputs []int
	for i := 0; i < 100000; i++ {
		outputs = append(outputs, race_condition())
	}
	print_shares([]int{-1, 0, 1}, outputs)
}
