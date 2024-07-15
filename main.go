package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env-local")
	if err != nil {
		fmt.Println("Error loading env file")
		return
	}
	// initialise new router
	r := chi.NewRouter()

	// add a middleware to log
	r.Use(middleware.Logger)

	// router functions
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	r.Get("/reverse/{data}", func(w http.ResponseWriter, r *http.Request) {
		str := chi.URLParam(r, "data")
		reversedStr := reverseString(str)

		w.Write([]byte(reversedStr))
	})
	r.Get("/mode", func(w http.ResponseWriter, r *http.Request) {
		mode := os.Getenv("MODE")

		w.Write([]byte(mode))
	})

	fmt.Println("server running in http://localhost:3000 🟢🚀")
	http.ListenAndServe(":3000", r)
	fmt.Println("server shutdown 🛑")
}

func reverseString(str string) string {
	runes := []rune(str) // Convert string to rune slice
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i] // Swap runes
	}
	return string(runes) // Convert rune slice back to string
}
