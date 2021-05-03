//Paquetes
package main

//Importaciones nativas e modulos externos
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//declaracion de tipos.
type tareas struct {
	ID       int    `json:ID`
	Nombre   string `json:Nombre`
	Apellido string `json:Apellido`
}

type alltareas []tareas

var task = alltareas{
	{
		ID:       1,
		Nombre:   "juan",
		Apellido: "ojeda",
	},
}

func gettareas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "aplication/json")
	json.NewEncoder(w).Encode(task)
}

func creartareas(w http.ResponseWriter, r *http.Request) {
	var newdata task
	body, err := ioutil(r.Body)

	if err != nil {
		fmt.Fprintf(w, "error dato invalido")
	}

	json.Unmarshal(body, &newdata)
	newdata.ID = len(task) + 1
	task = append(task, newdata)
	json.NewEncoder(w).Encode(task)
}

func inicio(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "bienvenido a mi api")
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", inicio)
	router.HandleFunc("/tareas", gettareas)
	server := http.ListenAndServe(":4000", router)
	log.Fatal(server)

}
