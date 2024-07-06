package main

import (
	"fmt"
	anb "internal/anbernicrc"
	"os"
)

func main() {
	for i, arg := range os.Args {
		if i == 0 {
			continue
		}
		crc, err := anb.FileCRC(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to get CRC of %s: %s", arg, err.Error())
			os.Exit(1)
		}
		fmt.Printf("%08x %s\n", crc, arg)
	}
}
