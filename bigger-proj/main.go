// modify the contents of an array with different go routines

package main

import (
	"fmt"
	"sync"
)

var m sync.Mutex

func main() {
	var wg sync.WaitGroup
	arr := []int{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			addVal(&arr, val)
		}(i)
	}
	wg.Wait()
	for _, i := range arr {
		fmt.Println(i)
	}
}

func addVal(arr *[]int, val int) {
	m.Lock()
	*arr = append(*arr, val)
	m.Unlock()
}
