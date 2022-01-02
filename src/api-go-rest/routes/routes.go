package routes

import (
	"log"
	"net/http"

	"github.com/BrunoIstvan/go-rest-api/controllers"
	"github.com/BrunoIstvan/go-rest-api/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {

	r := mux.NewRouter()
	r.Use(middleware.SetContentTypeIntoHeader)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalities", controllers.AllPersonalities).Methods("Get")
	r.HandleFunc("/api/personalities", controllers.CreateAPersonality).Methods("Post")
	r.HandleFunc("/api/personalities/{id}", controllers.DeleteAPersonality).Methods("Delete")
	r.HandleFunc("/api/personalities/{id}", controllers.UpdateAPersonality).Methods("Put")
	r.HandleFunc("/api/personalities/{id}", controllers.GetOnePersonality).Methods("Get")

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))

}
