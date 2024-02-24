package main

import (
	"fmt"
	"sync"
)

//

var wg sync.WaitGroup

func main() {
	for _, salutation := range []string{"Hello", "greeting", "good day"} {
		wg.Add(1)
		go func(salutation string) {
			defer wg.Done()
			fmt.Println(salutation)
		}(salutation)
	}
	wg.Wait()
}
