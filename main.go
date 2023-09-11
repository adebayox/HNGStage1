package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"time"
)

type Info struct {
	Slack_name string `json:"slackname"`
	Current_day string `json:"currentday"`
	Utc_time string `json:"utc_time"`
 	Track string `json:"track"` 
	Github_file_url string `json:"githubfileurl"`
	Github_repo_url string `json:"githubrepourl"`
	Status_code int `json:"statuscode"`
}

var infos []Info

func getInfos(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
	queryParams := r.URL.Query()
    slackName := queryParams.Get("slack_name")
    track := queryParams.Get("track")
	
	  // Initialize a slice to hold filtered results
	  var filteredInfos []Info

	  for _, item := range infos {
		  if item.Slack_name == slackName {
			  if item.Track == track {
				  filteredInfos = append(filteredInfos, item)
			  }
		  }
	  }

	  if err := json.NewEncoder(w).Encode(infos); err != nil {
		  http.Error(w, err.Error(), http.StatusInternalServerError)
		  return
	  }
}


func main() {
	r := mux.NewRouter()

	infos = append(infos, Info{
		Slack_name: "Adebayo David", 
		Current_day: time.Now().UTC().Weekday().String(), 
		Utc_time: time.Now().UTC().Format(time.RFC3339), 
		Track: "backend",
		Github_file_url: "www.github.com/adebayox", 
		Github_repo_url: "adebayox github", 
		Status_code: 200 })

	r.HandleFunc("/info", getInfos ).Methods("GET")

	fmt.Printf("starting server at 8000\n")
	log.Fatal(http.ListenAndServe(":8000",r))
}