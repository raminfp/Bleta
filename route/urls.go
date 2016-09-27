package route

import (
	"github.com/gorilla/mux"
	"github.com/raminfp/Bleta/app/controller"
	"net/http"
	"github.com/raminfp/Bleta/settings"

)

// URL Mapping
func Url() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	// serve files under http://localhost:8000/static/<filename>/,<otherfiles>
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(settings.StaticFilesDirs()))))

	// Define URL Pattern on `SAMIN WEB GO`
	// Example : router.HandleFunc("<your url >", `func view`).Methods("GET/PUT/POST")
	// Example Regex : router.HandleFunc("/user/{category}/{id:[0-9]+}",`funx view`).Methods("POST")

	router.HandleFunc("/", controller.IndexController).Methods("GET") // root path


	// Return ALL urls map
	return router
}




