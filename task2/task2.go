package main

import (
	"fmt"
)

func main() {

	n := [][]int{
		{2, 6, 9, 24},
		{7, 3, 94, 3, 0},
		{4, 2, 8, 35},
	}

	ch := make(chan int)
	done := make(chan bool)
	for _, v := range n {
		go func(slc []int) {
			result := 0
			for _, k := range slc {
				result += k
			}
			ch <- result
			done <- true
		}(v)
	}

	go func() {
		for range n {
			<-done
		}
		close(ch)
		close(done)
	}()

	for res := range ch {
		fmt.Println(res)
	}
}
