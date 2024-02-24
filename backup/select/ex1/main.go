package main

func main() {
	stopCh := make(chan struct{})
	requestCh := make(chan Request)
	resultCh := make(chan Result)

	go func() {
		for {
			var req Request
			select {
			case req = <-requestCh:
				// Received a reuqest to process
			case <-stopCh:
				// Stop requested, cleanup and return
				cleanup()
				return
			}
			// Do some processing
			someLongProcessing(req)
			// Check if stop requested another long task
			select {
			case <-stopCh:
				// Stop requested, cleanup and return
				cleanup()
				return
			default:
			}
			// Do more processing
			result := otherLongProcessing(req)
			select {
			// Wait until resultCh becomes sendable, or stop requested
			case resultCh <- result:
				// Result is set
			case <-stopCh:
				// Stop reqquested
				cleanup()
				return
			}
		}
	}()
}
