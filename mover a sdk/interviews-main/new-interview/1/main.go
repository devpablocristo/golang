// How to stop a goroutine

package main

func main() {
	quit := make(chan bool)
	go func() {
		for {
			select {
			case <-quit:
				return
			default:
				// …
			}
		}
	}()
	// …
	quit <- true
}
