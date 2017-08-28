// Package router initializes the application router
package router

import (
	// mux router
	"github.com/gorilla/mux"

	// Middleware chaining
	"github.com/justinas/alice"

	// Resources
	"github.com/krishamoud/game-server/app/bundles/containers"
	"github.com/krishamoud/game-server/app/bundles/users"

	// common middleware
	"github.com/krishamoud/game-server/app/common/middleware"

	"net/http"
)

// Router initializes and returns a mux.Router that will handle all api requests
func Router() *mux.Router {
	// Mux Router declaration
	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1/").Subrouter()

	// Controllers declaration
	cc := &containers.ContainersController{}
	uc := &users.UsersController{}

	// middleware chaining
	commonHandlers := alice.New(middleware.LoggingHandler, middleware.RecoverHandler)
	securedHandlers := commonHandlers.Append(middleware.Authenticate)

	// Container Information Routes
	s.Handle("/containers", securedHandlers.ThenFunc(cc.Index)).Methods("GET")
	s.Handle("/containers/{containerId}", securedHandlers.ThenFunc(cc.Show)).Methods("GET")

	// User creation Routes
	s.Handle("/users/new", commonHandlers.ThenFunc(uc.New)).Methods("GET")
	s.Handle("/users", commonHandlers.ThenFunc(uc.Create)).Methods("POST")

	// User Information Routes
	s.Handle("/users", securedHandlers.ThenFunc(uc.Index)).Methods("GET")
	s.Handle("/users/{userId}", securedHandlers.ThenFunc(uc.Show)).Methods("GET")

	// Auth Routes
	s.Handle("/auth", commonHandlers.ThenFunc(uc.Auth)).Methods("POST")

	// Naked route: only being used for testing purposes at the moment
	// change home.html to get logs for a certain container
	r.Handle("/", commonHandlers.ThenFunc(serveHome)).Methods("GET")

	return r
}

// serveHome returns home.html and is used for testing purposes only
func serveHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "home.html")
}