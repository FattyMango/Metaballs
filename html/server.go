package main

import (
	"fmt"
	"net/http"
)


func main() {
// make a server that serves metaballs.wasm with headers that allow it to be run in a browser
	http.Handle("/", Handler)
	http.ListenAndServe(":8080", nil)


}

// Handler is a http.Handler that serves metaballs.wasm with headers that allow it to be run in a browser
var Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Serving wasm")
	w.Header().Set("Content-Type", "application/wasm")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	http.ServeFile(w, r, "/home/kassab/GO/metaballs/html/metaballs.wasm")
})
