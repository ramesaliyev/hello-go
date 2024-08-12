package main

import (
	"fmt"
	"runtime"
)

func main() {
	defaultGOMAXPROCS := runtime.GOMAXPROCS(0)
	fmt.Printf("Default GOMAXPROCS: %d\n", defaultGOMAXPROCS)
}
