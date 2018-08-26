package main

import (
	"log"
	"os"
)

func init() {
	if cwd, err := os.Getwd(); err != nil {
		log.Fatalf("os.Getwd() failed: %v\n", err)
	} else {
		log.Printf("Current working dir: %s\n", cwd)
	}
}

func main() {
	if host, err := os.Hostname(); err != nil {
		log.Fatalf("os.Hostname() failed: %v\n", err)
	} else {
		log.Printf("Hostname: %s\n", host)
	}
}
