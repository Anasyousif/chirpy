package main

import ("log"
"net/http")

func main() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("."))
	mux.Handle("/app/",http.StripPrefix("/app",fileServer))

	mux.HandleFunc("/healthz", handlerReadiness)
	server :=&http.Server {
		Addr: ":8080",
		Handler: mux, 
	}

	log.Println("Serving on port: 8080")
	log.Fatal(server.ListenAndServe())
}

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}