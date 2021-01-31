package main

type user struct {
	Name     string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type kvSite struct {
}

type kvChannel struct {
}

type kvPlaylist struct{}
