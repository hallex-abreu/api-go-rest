package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/hallex-abreu/api-go-rest/controllers"
	"github.com/hallex-abreu/api-go-rest/middleware"
)

func HandleRequest() {
	route := mux.NewRouter()
	route.Use(middleware.ContentTypeMiddleware)
	route.HandleFunc("/", controllers.Home).Methods("Get")
	route.HandleFunc("/personalidades", controllers.Index).Methods("Get")
	route.HandleFunc("/personalidades/{id}", controllers.Show).Methods("Get")
	route.HandleFunc("/personalidades", controllers.Store).Methods("Post")
	route.HandleFunc("/personalidades/{id}", controllers.Update).Methods("Put")
	route.HandleFunc("/personalidades/{id}", controllers.Delete).Methods("Delete")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(route)))
}
