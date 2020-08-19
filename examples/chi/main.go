package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

func writeJSON(w http.ResponseWriter, obj interface{}) (err error) {
	buf, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(buf)
	return nil
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain;charset=UTF-8")
	w.Write([]byte("Hello,World!"))
}

type User struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

func userHandler(w http.ResponseWriter, _ *http.Request) {
	user := &User{
		ID:   123,
		Name: "Jhon",
	}
	writeJSON(w, user)
}

func main() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	}))

	router.Get("/", home)
	router.Get("/user", userHandler)

	server := http.Server{
		Handler: router,
		Addr:    ":80",
	}

	fmt.Println("Start server")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
}
