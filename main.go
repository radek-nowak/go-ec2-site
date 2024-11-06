package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type Pokemon struct {
	Name          string
	Type          string
	PokedexNumber string
	Level         int
	Weight        float64
}

func main() {

	testData := map[string][]Pokemon{
		"Pokemon": {
			{Name: "Pikachu", Type: "Electric", PokedexNumber: "025", Level: 15, Weight: 6.0},
			{Name: "Charmander", Type: "Fire", PokedexNumber: "004", Level: 10, Weight: 8.5},
			{Name: "Bulbasaur", Type: "Grass/Poison", PokedexNumber: "001", Level: 12, Weight: 6.9},
			{Name: "Squirtle", Type: "Water", PokedexNumber: "007", Level: 10, Weight: 9.0},
			{Name: "Jigglypuff", Type: "Normal/Fairy", PokedexNumber: "039", Level: 5, Weight: 5.5},
			{Name: "Meowth", Type: "Normal", PokedexNumber: "052", Level: 8, Weight: 4.2},
			{Name: "Psyduck", Type: "Water", PokedexNumber: "054", Level: 14, Weight: 19.6},
			{Name: "Machop", Type: "Fighting", PokedexNumber: "066", Level: 18, Weight: 19.5},
			{Name: "Geodude", Type: "Rock/Ground", PokedexNumber: "074", Level: 9, Weight: 20.0},
			{Name: "Eevee", Type: "Normal", PokedexNumber: "133", Level: 10, Weight: 6.5},
			{Name: "Snorlax", Type: "Normal", PokedexNumber: "143", Level: 30, Weight: 460.0},
			{Name: "Gengar", Type: "Ghost/Poison", PokedexNumber: "094", Level: 25, Weight: 40.5},
			{Name: "Onix", Type: "Rock/Ground", PokedexNumber: "095", Level: 20, Weight: 210.0},
			{Name: "Mewtwo", Type: "Psychic", PokedexNumber: "150", Level: 70, Weight: 122.0},
			{Name: "Gyarados", Type: "Water/Flying", PokedexNumber: "130", Level: 35, Weight: 235.0},
		},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templ := template.Must(template.ParseFiles("public/index.html"))
		templ.Execute(w, testData)
	})

	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("new-pokemon-name")
		pokemonType := r.FormValue("new-pokemon-type")
		pokedexNumber := r.FormValue("new-pokemon-number")
		weight, _ := strconv.ParseFloat(r.FormValue("new-pokemon-weight"), 32)
		level, _ := strconv.Atoi(r.FormValue("new-pokemon-level"))

		templ := template.Must(template.ParseFiles("public/index.html"))

		pokemon := Pokemon{
			Name:          name,
			Type:          pokemonType,
			PokedexNumber: pokedexNumber,
			Level:         level,
			Weight:        weight,
		}

		testData["Pokemon"] = append(testData["Pokemon"], pokemon)

		templ.ExecuteTemplate(w, "pokemon-list-element", pokemon)
	})

	log.Fatal(http.ListenAndServe(":3000", nil))

}
