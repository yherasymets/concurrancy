package main

import (
	"fmt"
	"sync"
)

func main() {
	n := [][]int{
		{2, 6, 9, 24},
		{7, 3, 94, 3, 0},
		{4, 2, 8, 35},
	}
	wg := sync.WaitGroup{}
	wg.Add(len(n))
	for i, v := range n {
		go func(i int, v []int) {
			defer wg.Done()
			sum(fmt.Sprintf("slice %d:", i), v)
		}(i, v)
	}
	wg.Wait()
}

func sum(name string, slc []int) {
	s := 0
	for _, v := range slc {
		s += v
	}
	fmt.Println(name, s)
}

