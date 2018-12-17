package ezio

import (
	"fmt"
	"time"
)

func Spinner(delay time.Duration, stopSpinner chan int) {
	for {
		select {
		default:
			for _, r := range `-\|/` {
				fmt.Printf("\rPlease wait: %c", r)
				time.Sleep(delay)
			}
		case <-stopSpinner:
			fmt.Printf("\r\rPlease wait: download complete")
			return

		}
	}
}
