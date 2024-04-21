package main

import (
	"log"
	"runtime"
)

func main() {
	// print server info
	log.Printf("******************************************")
	log.Printf("** %sCryptor Trend API%s v%s built in %s", "\033[31m", "\033[0m", "v1.0.0", runtime.Version())
	log.Printf("**----------------------------------------")
	log.Printf("** Running with %d Processors", runtime.NumCPU())
	log.Printf("** Running on %s", runtime.GOOS)
	log.Printf("******************************************")
}
