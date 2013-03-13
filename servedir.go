package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	endpt := flag.String("endpt", ":8080", "endpoint to listen on")
	dir := flag.String("dir", "./", "directory to serve--be careful!")
	flag.Parse()
	http.Handle("/", http.FileServer(http.Dir(*dir)))
	log.Print("Listening on ", *endpt)
	err := http.ListenAndServe(*endpt, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}

