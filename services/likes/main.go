package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

var likes = make(map[string][]string)

func getLikes(userId string) []string {
	userLikes := likes[userId]
	if userLikes == nil {
		userLikes = []string{}
	}
	return userLikes
}

func addLike(userId string, like string) {
	likes[userId] = append(getLikes(userId), like)
}

func getLikeFromRequest(request *http.Request) string {
	var body map[string]string
	json.NewDecoder(request.Body).Decode(&body)
	return body["like"]
}

func main() {
	router := chi.NewRouter()

	router.Get("/users/{id}/likes", func(response http.ResponseWriter, request *http.Request) {
		likes, _ := json.Marshal(map[string][]string{
			"likes": getLikes(chi.URLParam(request, "id")),
		})
		response.Write(likes)
	})

	router.Post("/users/{id}/likes", func(response http.ResponseWriter, request *http.Request) {
		addLike(chi.URLParam(request, "id"), getLikeFromRequest(request))
		response.WriteHeader(http.StatusCreated)
	})

	fmt.Println("Starting likes service")
	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", router)
}
