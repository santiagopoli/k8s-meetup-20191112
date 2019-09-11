package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func getLikes(userId string) map[string][]string {
	return map[string][]string{
		"likes": {"Kubernetes", "Open Policy Agent", "Cloud Native", "Docker"},
	}
}

func main() {
	router := chi.NewRouter()

	router.Get("/users/{id}/likes", func(response http.ResponseWriter, request *http.Request) {
		likes, _ := json.Marshal(getLikes(chi.URLParam(request, "id")))
		response.Write(likes)
	})

	fmt.Println("Starting likes service")
	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", router)
}
