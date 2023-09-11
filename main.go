package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Info struct {
    SlackName      string `json:"slackname"`
    CurrentDay     string `json:"currentday"`
    UTCTime        string `json:"utc_time"`
    Track          string `json:"track"`
    GithubFileURL  string `json:"githubfileurl"`
    GithubRepoURL  string `json:"githubrepourl"`
    StatusCode     int    `json:"statuscode"`
}


func getInfos(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    queryParams := r.URL.Query()
    slackName := queryParams.Get("slack_name")
    track := queryParams.Get("track")

    // Create an Info object based on the query parameters
    info := Info{
        SlackName:      slackName,
        CurrentDay:     time.Now().UTC().Weekday().String(),
        UTCTime:        time.Now().UTC().Format(time.RFC3339),
        Track:          track,
        GithubFileURL:  "https://github.com/adebayox/HNGStage1/blob/main/main.go",
        GithubRepoURL:  "https://github.com/adebayox/HNGStage1",
        StatusCode:     200,
    }

    // Encode the Info object as JSON and send it as the response
    if err := json.NewEncoder(w).Encode(info); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}



func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api", getInfos ).Methods("GET")

	cors := handlers.CORS(
        handlers.AllowedOrigins([]string{"*"}), // Replace with your specific allowed origins
        handlers.AllowedMethods([]string{"GET", "OPTIONS"}), // Add other HTTP methods if needed
        handlers.AllowedHeaders([]string{"Content-Type"}),
    )

    // Wrap your router with the CORS middleware
    http.Handle("/", cors(r))

	fmt.Printf("starting server at 8000\n")
	log.Fatal(http.ListenAndServe(":8000",r))
}