package main

import (
	"encoding/json"
	"fmt"
	"movies-be/internal/models"
	"net/http"
	"time"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "Go Backend Up and Running",
		Version: "1.0.0",
	}

	out, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func (app *application) AllMovies(w http.ResponseWriter, r *http.Request) {
	var movies []models.Movie

	rd, _ := time.Parse("2006-01-02", "2013-10-05")

	walterMitty := models.Movie{
		ID:          1,
		Title:       "The Secret Life of Walter Mitty",
		ReleaseDate: rd,
		RunTime:     114,
		Rated:       "PG",
		Description: "When both he and a colleague are about to lose their job, Walter takes action by embarking on an adventure more extraordinary than anything he ever imagined.",
	}

	movies = append(movies, walterMitty)

	rd, _ = time.Parse("2006-01-02", "1995-09-22")

	empireRecords := models.Movie{

		ID:          2,
		Title:       "Empire Records",
		ReleaseDate: rd,
		RunTime:     90,
		Rated:       "PG-13",
		Description: "A day in the life of the employees of Empire Records. Except this is a day where everything comes to a head for a number of them facing personal crises - can they pull through together? And more importantly, can they keep their record store independent and not swallowed up by corporate greed?",
	}

	movies = append(movies, empireRecords)

	out, err := json.Marshal(movies)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
