//https://www.digitalocean.com/community/tutorials/how-to-make-an-http-server-in-go
//https://blog.logrocket.com/creating-a-web-server-with-golang/

package server

import (
	"fmt"
	"net/http"
	"log"
)

func Run() {
    fmt.Println("The server has started successfully");
	// file handle
	fileServer := http.FileServer(http.Dir("./public"))
    http.Handle("/", fileServer)

	if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}