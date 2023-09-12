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
	SlackName  string
	CurrentDay string
	UTCTime    string
	Track      string
	GithubFile string
	GithubRepo string
	StatusCode int
}

func ParseBody(r *http.Request, X interface{}) {
	slack_name := "ValGrace"
	current_day := time.Now().Weekday()
	current_time := time.Now()
	utc_time := current_time.UTC()
	utcTime := utc_time.Format(time.RFC3339Nano)
	track := "backend"
	githubFile := "https://github.com/ValGrace/go-hngproject/main.go"
	githubRepo := "https://github.com/ValGrace/go-hngproject"
	profileData := fmt.Sprintf(`{"SlackName": "%s", "CurrentDay": "%s", "UTCTime": "%s", "Track": "%s", "GithubFile": "%s", "GithubRepo": "%s", "StatusCode": "%v"}`, slack_name, current_day, utcTime, track, githubFile, githubRepo, http.StatusCreated)
	reader := []byte(profileData)
	json.Unmarshal(reader, X)
	return
}
func GetProfileById(r *http.Request) *Profile {
	var getProfile Profile
	slack_profile := r.URL.Query().Get("slack_name")
	track := r.URL.Query().Get("track")

	if slack_profile == "ValGrace" && track == "backend" {
		return &getProfile
	}

}
func GetProfile(w http.ResponseWriter, r *http.Request) {
	//   vars := mux.Vars(r)
	profileDetails := GetProfileById(r)
	res, _ := json.Marshal(profileDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

var RegisterGetRoute = func(router *mux.Router) {
	router.HandleFunc("/profile/", GetProfile).Methods("GET")
}

func main() {
	r := mux.NewRouter()

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
