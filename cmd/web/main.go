package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	// define a command line flag with a name "addr" and default value :4000
	addr := flag.String("addr", ":4000", "server address port")
	// parse the command line flag to get the value of addr
	flag.Parse()
	// add leveled logging capability i.e info and error
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}
	infoLog.Printf("Starting server at port %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
