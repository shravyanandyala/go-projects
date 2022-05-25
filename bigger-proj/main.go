// modify the contents of an array with different go routines

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var m sync.Mutex

func main() {
	var wg sync.WaitGroup
	s := []int{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			addVal(&s, val)
		}(i)
	}
	wg.Wait()
	fmt.Println()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			removeVal(&s)
		}(i)
	}
	wg.Wait()
}

func addVal(s *[]int, val int) {
	m.Lock()
	*s = append(*s, val)
	fmt.Println(val)
	m.Unlock()
}

func removeVal(s *[]int) {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	m.Lock()
	d := r.Intn(len(*s))
	fmt.Println((*s)[d])
	*s = append((*s)[0:d], (*s)[d+1:]...)
	m.Unlock()
}
