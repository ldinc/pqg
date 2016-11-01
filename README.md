Golang simple cyclic group generator Zp with p = 2q + 1, p,q - safe primes.

n - bit width for big.Int.

Based on Miller-Rabin test.

## Example:

```go
package main

import (
	"flag"
	"fmt"
	"github.com/ldinc/pqg"
	"log"
	"math/big"
	"sync"
	"time"
)

func timeTrack(start time.Time, msg string) {
	elapsed := time.Since(start)
	log.Printf("> %s: %s\n", msg, elapsed)
}

func main() {
	prop := flag.Int("p", 64, "Miller-Rabin test with 1 - 1/(4^p)")
	n := flag.Int("n", 4, "number of generated <p,q,g>")
	bits := flag.Int("bits", 1024, "number of generated <p,q,g>")
	flag.Parse()
	mutex := &sync.Mutex{}
	result := make(chan *big.Int)
	for i := 0; i < *n; i++ {
		go func(id int) {
			defer timeTrack(time.Now(), fmt.Sprintf("gen %d", id))
			p, q, g, err := pqg.Gen(*bits, *prop)
			if err != nil {
				panic(err)
			}
			mutex.Lock()
			result <- p
			result <- q
			result <- g
			mutex.Unlock()
		}(i)
	}
	for i := 0; i < *n; i++ {
		log.Println("p = ", <-result)
		log.Println("q = ", <-result)
		log.Println("g = ", <-result)
	}
}
```
