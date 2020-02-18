package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"plugin"
)

func main() {
	pluginName := flag.String("plugin", "plugin.goos", "Plugin file name")
	flag.Parse()

	p, err := plugin.Open(*pluginName)
	if err != nil {
		log.Fatalf("failed to open plugin, %v, %v", *pluginName, err)
	}

	v, err := p.Lookup("Random")
	if err != nil {
		log.Fatalf("failed to lookup Random, %v", err)
	}
	randomFn = v.(func([]byte) (int, error))

	b := make([]byte, 24)
	n, err := randomFn(b)
	if err != nil {
		log.Fatalf("failed to get random bytes, %v", err)
	}

	fmt.Println("Got bytes, ", n)
	fmt.Println(hex.Dump(b[:n]))
}

var randomFn func([]byte) (int, error)
