package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Profile struct {
	SlackName  string `json:"slack_name"`
	CurrentDay string `json:"current_day"`
	UTCTime    string `json:"utc_time"`
	Track      string `json:"track"`
	GithubFile string `json:"github_file_url"`
	GithubRepo string `json:"github_repo_url"`
	StatusCode int    `json:"status_code"`
}

func GetProfile(w http.ResponseWriter, r *http.Request) {
	slack := r.URL.Query().Get("slack_name")
	trackurl := r.URL.Query().Get("track")

	if slack == "" || trackurl == "" {
		http.Error(w, "Both track and slack name parameters are reuired ", http.StatusBadRequest)
		return
	}
	slack_name := "ValGrace"
	current_day := time.Now().Weekday()
	current_time := time.Now()
	utc_time := current_time.UTC()
	utcTime := utc_time.Format(time.RFC3339Nano)
	track := "backend"
	githubFile := "https://github.com/ValGrace/go-hngproject/main.go"
	githubRepo := "https://github.com/ValGrace/go-hngproject"

	profileData := Profile{
		SlackName:  slack_name,
		CurrentDay: current_day.String(),
		UTCTime:    utcTime,
		Track:      track,
		GithubFile: githubFile,
		GithubRepo: githubRepo,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(profileData)

}

var RegisterGetRoute = func(router *mux.Router) {
	router.HandleFunc("/profile", GetProfile).Methods("GET")
}

func main() {
	r := mux.NewRouter()
	RegisterGetRoute(r)
	fmt.Printf("starting server at port 8080")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
