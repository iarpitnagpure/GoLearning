package main

import (
	"fmt"
	"sync"
)

// Can also create mutext seperately with var like below
// var (
//     counter int
//     mu      sync.Mutex
//     wg      sync.WaitGroup
// )

type post struct {
	view int
	mu   sync.Mutex
}

func (p *post) incrementView(wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
		p.mu.Unlock()
	}()
	p.mu.Lock()
	p.view += 1
}

// Mutex (mutual exclusion lock) is used to prevent multiple goroutines from accessing shared data at the same time — it ensures safety in concurrent code.
func main() {
	var wg sync.WaitGroup
	newPost := post{
		view: 0,
	}

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go newPost.incrementView(&wg)
	}

	wg.Wait()
	fmt.Println("newPost", newPost.view)
	// Output without mutex: "newPost --->" 9758 or 9634
	// Getting everytime new output since multiple goroutines shared data at same time
	// Output with mutex: "newPost --->" 10000 everytime
}
