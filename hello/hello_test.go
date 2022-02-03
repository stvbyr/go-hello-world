package hello_test

import (
	"strings"
	"testing"

	"stvbyr.tech/go/hello-world/hello"
)

func TestBestHelloWorldString(t *testing.T) {
	if !strings.HasPrefix(hello.BestHelloWorldString(), "Hello") {
		t.Fatal("Hello World must begin with 'Hello'")
	}
}
