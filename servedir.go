package main

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"
)

type spaHandler struct {
  root http.FileSystem
  index string
}

func SpaHandler(root http.FileSystem) http.Handler {
  return &spaHandler{root, "index.html"}
}

func (h *spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  path := filepath.Clean(r.URL.Path)
  if path == "" || path == "/" {
    path = h.index
  }
  f, err := h.root.Open("/" + path)
  if err != nil {
    f, err = h.root.Open("/" + h.index)
  }
  if err != nil {
    http.NotFound(w, r)
    return
  }
  defer f.Close()

  i, err := f.Stat()
  if err != nil || i.IsDir() {
    http.NotFound(w, r)
    return
  }
  http.ServeContent(w, r, i.Name(), i.ModTime(), f)
}

func main() {
	endpt := flag.String("endpt", ":8080", "endpoint to listen on")
	dir := flag.String("dir", "./", "directory to serve--be careful!")
	spa := flag.Bool("spa", false, "try to serve index.html instead of 404s")
	flag.Parse()

  root := http.Dir(*dir)
	if *spa {
		http.Handle("/", SpaHandler(root))
	} else {
		http.Handle("/", http.FileServer(root))
	}
	log.Print("Listening on ", *endpt)
	err := http.ListenAndServe(*endpt, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}

