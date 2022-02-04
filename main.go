package main

import (
	"fmt"

	"rsc.io/quote"
	"stvbyr.tech/go/hello-world/hello"
	"stvbyr.tech/go/hello-world/httpserver"
)

func main() {
	fmt.Println(hello.BestHelloWorldString())
	fmt.Println(quote.Go())

	httpserver.CreateHttpServer()
}
