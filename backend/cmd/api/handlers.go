package main

import (
	"net/http"

	"github.com/hiroshi-iwashita/20230207practice_docker-nuxt-go-ecs/internal/data"
)

// jsonResponse is the type used for generic JSON responses
type jsonResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type envelope map[string]interface{}

func (app *application) test(w http.ResponseWriter, r *http.Request) {
	// io.WriteString(w, "Hello world")
	payload := jsonResponse{
		Error:   false,
		Message: "success",
		Data:    "Hello go",
	}

	err := app.writeJSON(w, http.StatusOK, payload)
	if err != nil {
		app.errorLog.Println(err)
	}
}

func (app *application) AllUsers(w http.ResponseWriter, r *http.Request) {
	var users data.User
	all, err := users.GetAll()
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "success",
		Data:    envelope{"users": all},
	}

	app.writeJSON(w, http.StatusOK, payload)
}
