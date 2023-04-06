package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "Welcome to my Players API",
		Version: "1.0.0",
	}
	out, err := json.Marshal(payload)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(out)
	if err != nil {
		log.Println(err)
	}
}

func (app *application) GetPlayers(w http.ResponseWriter, r *http.Request) {
	players, err := app.DB.GetPlayers()
	if err != nil {
		log.Println(err)
	}
	out, err := json.Marshal(players)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(out)
	if err != nil {
		log.Println(err)
	}
}

func (app *application) OnePlayer(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	playerID, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("Error:", err)
		w.WriteHeader(http.StatusBadRequest)
	}
	player, err := app.DB.OnePlayer(playerID)
	if err != nil {
		log.Println(err)
	}
	out, err := json.Marshal(player)
	if err != nil {
		fmt.Println("Error:", err)
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(out)
	if err != nil {
		fmt.Println("Error:", err)
		w.WriteHeader(http.StatusBadRequest)
	}

}
