package handler

import (
	"fmt"
	"net/http"

	"github.com/KarthikUCH/buildapi/db"
	"github.com/gorilla/mux"
)

var dbInstance db.Database

func Router(db db.Database) *mux.Router {
	router := mux.NewRouter().StrictSlash(true).UseEncodedPath()
	dbInstance = db

	subRouterV3 := router.PathPrefix("/v3/").Subrouter().StrictSlash(true).UseEncodedPath()

	subRouterV3.NotFoundHandler = http.HandlerFunc(methodNotFoundHandler)

	// routing
	subRouterV3.HandleFunc("/", servehome).Methods(http.MethodGet)

	subRouterV3.HandleFunc("/items", GetItems(dbInstance)).Methods(http.MethodGet)
	subRouterV3.HandleFunc("/item", SaveItem(dbInstance)).Methods(http.MethodPost)
	subRouterV3.HandleFunc("/item", GetItem(dbInstance)).Methods(http.MethodGet)
	subRouterV3.HandleFunc("/item", UpdateItem(dbInstance)).Methods(http.MethodPut)
	subRouterV3.HandleFunc("/item", DeleteItem(dbInstance)).Methods(http.MethodDelete)

	return router
}

func methodNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	fmt.Print(w, "Not Found! Please check again.")
}

// serve home route
func servehome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by Raju Karthikeyan</h1>"))
}
