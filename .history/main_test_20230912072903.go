package main

import (
	"testing"
	"net/http/httptest"
	"time"
)

func GetUserTest(t *testing.T) {
	type testUser struct {
		name       string
		method     string
		profile *Profiles
		want string
		statusCode int
	}{
		{
			name: "profile exists",
			method: http.MethodGet,
			profile: &Profiles{
				Profile{
					slackname: "ValGrace",
					current_day: time.Now().Weekday(),
					current_time: "2023-9"
					track: "backend",
                    github_file_url: "https://github.com/ValGrace/go-hngproject/main.go",
					github_repo_url: "https://github.com/ValGrace/go-hngproject",
					statusCode: http.StatusOK   
				}
			} 
		}
	}
}

func TestGetMethod(t *testing.T) {
	
}
