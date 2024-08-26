package main

import "net/http"

func main() {
	server := httpServer{}

	http.ListenAndServe(":8090", server)
}

type httpServer struct {
}

func (h httpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Write([]byte(`Hello World`))
	}
}
