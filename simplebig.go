package main

import (
	"fmt"
	"simplebig/simplefloat"
)

func main() {
	f := simplefloat.New(12.25)
	fmt.Println(f.Float64())
	fmt.Println(f.Quo(simplefloat.New(5)).Float64())
}
