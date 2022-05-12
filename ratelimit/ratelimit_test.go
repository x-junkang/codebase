package ratelimit

import (
	"fmt"
	"sync"
	"testing"
)

func TestPass(t *testing.T) {
	limiter := New(5, 20)
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			if limiter.Pass() {
				fmt.Println("pass")
			} else {
				fmt.Println("failed")
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
