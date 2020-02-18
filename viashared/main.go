package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"log"
)

func main() {
	atOnce := flag.Int("p", 1, "Number of C threads to run at once")
	calls := flag.Int("c", 100, "Number of C -> Go callbacks to run per C thread")
	flag.Parse()

	b := make([]byte, 24)
	n, err := Random(b)
	if err != nil {
		log.Fatalf("failed to get random bytes, %v", err)
	}

	fmt.Println("Got bytes, ", n)
	fmt.Println(hex.Dump(b[:n]))

	runCallbacks(*atOnce, *calls)
}
