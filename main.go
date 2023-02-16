package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Note struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

var notes []Note

func main() {
	r := mux.NewRouter()

	appendData()

	r.HandleFunc("/", welcome)
	r.HandleFunc("/notes", getNotes).Methods("GET")
	r.HandleFunc("/notes/{id}", getNote).Methods("GET")
	r.HandleFunc("/notes", createNote).Methods("POST")
	r.HandleFunc("/notes/{id}", deleteNote).Methods("DELETE")
	r.HandleFunc("/notes/{id}", updateNote).Methods("PUT")

	fmt.Printf("Starting Server\n")

	log.Fatal(http.ListenAndServe(":8080", r))

}

func welcome(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to Notes API"))
}

func updateNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range notes {
		if item.Id == params["id"] {
			//delete item at id pos
			notes = append(notes[:index], notes[index+1:]...)

			//update item at id pos
			var note Note
			_ = json.NewDecoder(r.Body).Decode(&note)
			note.Id = params["id"]
			notes = append(notes, note)
			json.NewEncoder(w).Encode(note)
			fmt.Printf("Note Updated\n")
			return
		}
	}
}

func createNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var note Note
	_ = json.NewDecoder(r.Body).Decode(&note)
	notes = append(notes, note)
	fmt.Printf("Note Created\n")
	json.NewEncoder(w).Encode(note)
}

func getNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range notes {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func deleteNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range notes {
		if item.Id == params["id"] {
			notes = append(notes[:index], notes[(index+1):]...)
			fmt.Printf("Note Deleted\n")
			break
		}
	}
	json.NewEncoder(w).Encode(notes)
}

func getNotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notes)
}

func appendData() {
	notes = append(notes, Note{
		Id:          "1",
		Title:       "First Note",
		Description: "This is the first note.",
	})

	notes = append(notes, Note{
		Id:          "2",
		Title:       "Second Note",
		Description: "This is the second note.",
	})

	notes = append(notes, Note{
		Id:          "3",
		Title:       "Third Note",
		Description: "This is the third note.",
	})
}
