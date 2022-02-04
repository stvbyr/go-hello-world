package httpserver

import (
	"io"
	"log"
	"net/http"
	"os"

	"stvbyr.tech/go/hello-world/readfile"
)

func CreateHttpServer() {
	http.HandleFunc("/", handler)

	port := ":8080"
	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func handler(writer http.ResponseWriter, r *http.Request) {
	path, _ := os.Getwd()
	contents, err := readfile.ReadMarkdownFile(path + "/Hello.md")
	if err != nil {
		log.Fatal(err)
	}
	io.WriteString(writer, contents)
}
