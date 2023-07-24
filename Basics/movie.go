package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"Humphrey Bogart", "Ingrid bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true, Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true, Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
}

func main() {
	fmt.Print(movies)
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatal("JSON marshalling failed: %s", err)
	}
	fmt.Printf("\n%s\n", data)
	data, err0 := json.MarshalIndent(movies, "", "   ")
	if err0 != nil {
		log.Fatal("JSON marshalling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	// unmarshal to get only titles
	var titles []struct{ Title string }
	if err1 := json.Unmarshal(data, &titles); err1 != nil {
		log.Fatal("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(titles)
}
