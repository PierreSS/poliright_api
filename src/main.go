package main

//Ma librairie
import (
	"fmt"
	"time"
)

var (
	Version = "1.0.0"
	Build   = time.Now()
)

func main() {
	fmt.Printf("%s\n%s\n", Build, Version)
}
