//https://www.digitalocean.com/community/tutorials/how-to-make-an-http-server-in-go
//https://blog.logrocket.com/creating-a-web-server-with-golang/

package main

import (
	"fmt"
	"net/http"
	"log"
)

func main() {
    fmt.Println("Hello, World!");
	// file handle
	fileServer := http.FileServer(http.Dir("./public"))
    http.Handle("/", fileServer)
	
	http.HandleFunc("/form", formHandler)
	//
	if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}

func formHandler(w http.ResponseWriter, r *http.Request) {
    if err := r.ParseForm(); err != nil {
        fmt.Fprintf(w, "ParseForm() err: %v", err)
        return
    }
    fmt.Fprintf(w, "POST request successful")
    name := r.FormValue("name")
    address := r.FormValue("address")

    fmt.Fprintf(w, "Name = %s\n", name)
    fmt.Fprintf(w, "Address = %s\n", address)
}