package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/igorcavalcantedb/golang-rest-crud/controllers"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.PersonalidadeById).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.CriarNovaPersonalidade).Methods("Post")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletarPersonalidadeById).Methods("Delete")
	r.HandleFunc("/api/personalidades/{id}", controllers.EditarPersonalidadeById).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", r))
}
