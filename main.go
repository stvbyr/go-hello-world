package main

import (
	"fmt"

	"rsc.io/quote"
	"stvbyr.tech/go/hello-world/hello"
)

func main() {
	fmt.Println(hello.BestHelloWorldString())
	fmt.Println(quote.Go())
}
