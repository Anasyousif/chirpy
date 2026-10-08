package main

import ("log"
"net/http")

func main() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("."))
	mux.Handle("/",fileServer)
	server :=&http.Server {
		Addr: ":8080",
		Handler: mux, 
	}

	log.Println("Serving on port: 8080")
	log.Fatal(server.ListenAndServe())
}