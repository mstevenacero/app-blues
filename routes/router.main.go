package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	users "github.com/mstevenacero/application-core-blues/routes/users"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome blues core --->"))
}

func RegisterRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler).Methods("GET")
	router.HandleFunc("/user", users.GetUsersRoute).Methods("GET")
	router.HandleFunc("/user", users.AddUsersRoute).Methods("POST")
	return router
}
