package ezio

import (
	"fmt"
	"time"
)

func Spinner(delay time.Duration, quit chan int) {
	for {
		select {
		default:
			for _, r := range `-\|/` {
				fmt.Printf("\rPlease wait %c", r)
				time.Sleep(delay)
			}
		case <-quit:
			fmt.Printf("\r")
			return

		}
	}
}
