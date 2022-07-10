package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	infoLog  *log.Logger
	errorLog *log.Logger
}

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT Must Be Set")
	}

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		infoLog:  infoLog,
		errorLog: errorLog,
	}

	flag.Parse()

	srv := &http.Server{
		Addr:     *port,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Listening On Port %s", *port)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)

}
