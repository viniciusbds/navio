package loader

import (
	"fmt"
	"sync"
	"time"
)

// Loader displays an animation while waiting for a job to finish
func Loader(finishMsg string, done chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

	i := 0
	load := []rune(`|\-/`)

	for {
		select {
		case <-done:
			fmt.Printf("\r")
			fmt.Println(finishMsg)
			return
		default:
			fmt.Printf("\r")
			fmt.Printf(string(load[i]))
			time.Sleep(time.Millisecond * 100)
			i++
			if i == len(load) {
				i = 0
			}
		}
	}
}