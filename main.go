package main

import (
	"log"
	"net/http"
	"os"
	"encoding/json"
	"github.com/joho/godotenv"
	"github.com/gorilla/mux"
)

type Data struct {
	Message string `json:"message"`
	Data Pokemon `json:"data"`
}

type Message struct {
	Message string `json:"message"`
}

type Pokemon struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Weight int `json:"weight"`
}

func main() {
	godotenv.Load()
	handleRequests()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":" + port, nil); err != nil {
		log.Fatal(err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Message{Message: "Use /pokemon/{name} or /pokemon/{pokedexNumber} to request pokemon data"})
}

func getPokemon(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	pokemon := new(Pokemon)

	response, err := http.Get("https://pokeapi.co/api/v2/pokemon/" + vars["param"])

    if err != nil {
		log.Fatal(err)
    }

	json.NewDecoder(response.Body).Decode(pokemon)
	
	if pokemon.Id == 0 {
		json.NewEncoder(w).Encode(Message{Message: "Pokemon not Found"})
	} else {
		json.NewEncoder(w).Encode(Data{Message: "Ok", Data: *pokemon})
	}

}

func handleRequests() {
	r := mux.NewRouter()
	r.HandleFunc("/pokemon/{param}", getPokemon)
	r.HandleFunc("/", home)
	http.Handle("/", r)

}
