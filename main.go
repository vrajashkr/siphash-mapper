package main

import (
	"fmt"
	"os"

	"github.com/dchest/siphash"
	"github.com/pelletier/go-toml/v2"
)

type Input struct {
	HashKey string
	Items   []string
	Targets []string
}

func main() {
	configContents, err := os.ReadFile("./config.toml")
	if err != nil {
		fmt.Printf("[ERR] failed to read config file contents. error=%s\n", err)
		os.Exit(1)
	}

	var input Input

	err = toml.Unmarshal(configContents, &input)
	if err != nil {
		fmt.Printf("[ERR] failed to unmarshal TOML contents. error=%s\n", err)
		os.Exit(1)
	}

	numTargets := uint64(len(input.Targets))

	for _, item := range input.Items {
		h := siphash.New([]byte(input.HashKey))
		h.Write([]byte(item))
		sum64 := h.Sum64()
		indexOfTarget := sum64 % numTargets
		fmt.Printf("%s mapped to %s at index %d\n", item, input.Targets[indexOfTarget], indexOfTarget)
	}
}
