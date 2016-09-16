package main

import "net/http"
import "os"

func main() {
    http.Handle("/data/", http.StripPrefix("/data/", http.FileServer(http.Dir(os.Getenv("SNAP_DATA") + "/json"))))
    http.Handle("/", http.FileServer(http.Dir(os.Getenv("SNAP") + "/public_html")))
    http.ListenAndServe(":5080", nil)
}

