package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

type Handler struct {
	mux *http.ServeMux
}

func (p *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.mux.ServeHTTP(w, r)
	fmt.Printf(
		"%s %s %s\n",
		time.Now().Format(time.RFC3339),
		r.Method,
		r.URL.String(),
	)
}

func NewHandler() *Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/plain;charset=UTF-8")
		w.WriteHeader(200)
		_, _ = w.Write([]byte("Hello,World!"))
	})
	return &Handler{
		mux: mux,
	}
}

func main() {
	handler := NewHandler()

	server := http.Server{
		Handler: handler,
		Addr:    ":80",
	}

	fmt.Println("start server")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
