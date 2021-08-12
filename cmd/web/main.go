package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	// define a command line flag with a name "addr" and default value :4000
	addr := flag.String("addr", ":4000", "server address port")
	// parse the command line flag to get the value of addr
	flag.Parse()
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	log.Printf("Starting server at port %s", *addr)
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
