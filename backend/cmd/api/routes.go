package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/hiroshi-iwashita/20230207practice_docker-nuxt-go-ecs/internal/data"
)

// routes generates our routes and attaches them to handlers, using the chi router
// note that we return type http.Handler, and not *chi.Mux; since chi.Mux satisfies
// the interface requirements for http.Handler, it makes sense to return the type
// that is part of the standard library.
func (app *application) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{
			"https://*",
			"http://*",
		},
		AllowedMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"OPTIONS",
		},
		AllowedHeaders: []string{
			"Accept",
			"Authorization",
			"Content-Type",
			"X-CSRF-Token",
		},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	
	mux.Get("/", app.test)
	mux.Get("/users", app.AllUsers)
	mux.Post("/users/login", app.AllUsers)

	mux.Post("/users/get/{id}", app.GetUser)

	mux.Get("/users/add", func(w http.ResponseWriter, r *http.Request){
		var user data.User

		var u = data.User{
			Email: "you@there.com",
			FirstName: "You",
			LastName: "There",
			Password: "passoword",
		}

		app.infoLog.Println("Adding user...")

		id, err := user.Insert(u)
		if err != nil {
			app.errorLog.Println(err)
			app.errorJSON(w, err, http.StatusForbidden)
		}

		app.infoLog.Println("Got back id of", id)
		newUser, _ := user.GetOne(id)
		app.writeJSON(w, http.StatusOK, newUser)
	})


	return mux
}
