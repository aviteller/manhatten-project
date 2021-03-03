package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

const sessionLength int = 30

var dbSessions = map[string]session{}

func init() {
	initDB()
}

func signup(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var u user
	err := decoder.Decode(&u)

	if err != nil {
		panic(err)
	}

	if u.Name == "" || u.Email == "" || u.Password == "" {
		http.Error(w, "PLEASE ENTER ALL FIELDS", http.StatusBadRequest)
		return
	}

	if val, ok := dbUsers[u.Name]; ok {
		http.Error(w, "User: "+val.Name+" Already added", http.StatusBadRequest)
		return

	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u.Password = string(hashedPassword)
	newDBUserID := len(dbUsers) + 1
	u.ID = newDBUserID
	dbUsers[u.Name] = u

	fmt.Fprintln(w, "User added", u, dbUsers)
}

func login(w http.ResponseWriter, r *http.Request) {

	if alreadyLoggedIn(w, r) {
		http.Error(w, "Already logged in", http.StatusBadRequest)
		return

	}
	decoder := json.NewDecoder(r.Body)
	var u user
	err := decoder.Decode(&u)

	if err != nil {
		panic(err)
	}

	if u.Name == "" || u.Password == "" {
		http.Error(w, "PLEASE ENTER ALL FIELDS", http.StatusBadRequest)
		return
	}

	if foundUser, ok := dbUsers[u.Name]; !ok {
		http.Error(w, "User: "+foundUser.Name+" Not Found", http.StatusBadRequest)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUsers[u.Name].Password), []byte(u.Password))

	if err != nil {
		http.Error(w, "passwords do not match", http.StatusBadRequest)
		return
	}

	u1 := uuid.Must(uuid.NewV4())
	expiration := time.Now().Add(365 * 24 * time.Hour)

	cookie := &http.Cookie{
		Name:    "man-session-id",
		Value:   u1.String(),
		Expires: expiration,
	}

	http.SetCookie(w, cookie)
	dbSessions[cookie.Value] = session{u.Name, time.Now()}

	// fmt.Println(dbSessions)

	res := Message(true, "success")
	res["data"] = u
	Respond(w, res)
	return

}

func alreadyLoggedIn(w http.ResponseWriter, r *http.Request) bool {
	c, err := r.Cookie("man-session-id")

	if err != nil {
		return false
	}
	s, ok := dbSessions[c.Value]
	if ok {
		s.lastActive = time.Now()
		dbSessions[c.Value] = s
	}
	_, ok = dbUsers[s.un]
	// refresh session
	c.MaxAge = sessionLength
	http.SetCookie(w, c)
	// fmt.Println(dbUsers[s.un])
	return ok
}

func greet(w http.ResponseWriter, r *http.Request) {

	c, err := r.Cookie("man-session-id")

	if err != nil {
		panic(err)
	}

	// fmt.Println(c)

	if c != nil {

		var u user

		if val, ok := dbSessions[c.Value]; ok {

			u = dbUsers[val.un]

			fmt.Fprintln(w, "Hello! ", u.Name, "With Email ", u.Email)
			return
		}
	}

	// fmt.Println("Hello World! %s", time.Now())
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func getUser(w http.ResponseWriter, r *http.Request) user {
	var u user
	c, err := r.Cookie("man-session-id")

	if err != nil {
		panic(err)
	}

	if c != nil {

		if val, ok := dbSessions[c.Value]; ok {

			u = dbUsers[val.un]

			return u
		}
	}
	return u
}

func getChannels(siteID int) []kvChannel {
	var returnChannels []kvChannel

	for _, channel := range dbChannels {
		if siteID == channel.SiteID {

			playlists := getChannelPlaylists(channel.ID)

			fullChannel := kvChannel{
				channel.ID,
				channel.SiteID,
				channel.Name,
				playlists,
			}

			returnChannels = append(returnChannels, fullChannel)
		}
	}

	return returnChannels
}

func getChannelPlaylists(channelID int) []kvPlaylist {
	var kvPlaylists []kvPlaylist

	for _, link := range dbPlaylistChannelLinks {

		if link.ChannelID == channelID {

			for _, playlist := range dbPlaylists {
				if link.PlaylistID == playlist.ID {
					tracks := getPlaylistTracks(playlist.ID)

					fullPlaylist := kvPlaylist{
						playlist.ID,
						playlist.Name,
						tracks,
					}

					kvPlaylists = append(kvPlaylists, fullPlaylist)
				}
			}
		}
	}

	return kvPlaylists
}

func getPlaylistTracks(playlistID int) []kvTrack {
	var kvTracks []kvTrack

	for _, track := range dbTracks {

		if track.PlayistID == playlistID {

			kvTracks = append(kvTracks, track)

		}
	}

	return kvTracks
}
func getUserSites(w http.ResponseWriter, r *http.Request) {
	u := getUser(w, r)

	var returnSites []kvSite
	for _, site := range dbSites {
		if site.UserID == u.ID {

			chanls := getChannels(site.ID)

			fullSite := kvSite{
				site.ID,
				site.UserID,
				site.Name,
				chanls,
			}

			returnSites = append(returnSites, fullSite)
		}
	}
	res := Message(true, "success")

	res["data"] = returnSites

	Respond(w, res)
	// json.NewEncoder(w).Encode(returnSites)
	// getChannelPlaylists(re)
	// fmt.Fprintln(w, returnSites)
	// return returnSites
}

func logout(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("man-session-id")

	if err != nil {
		panic(err)
	}

	if c != nil {

		delete(dbSessions, c.Value)

	}

}

func logedin(w http.ResponseWriter, r *http.Request) {
	user := getUser(w, r)
	if user.ID != 0 {

		res := Message(true, "success")
		res["data"] = user
		Respond(w, res)
		return
	}
	res := Message(false, "fail")
	Respond(w, res)
}

func router(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	method := r.Method
	path := r.URL.Path
	path = strings.Replace(path, "/", "", -1)

	switch method {
	case "GET":
		switch path {
		case "greet":
			greet(w, r)
		case "logedin":
			logedin(w, r)
		case "logout":
			logout(w, r)
		case "sites":
			getUserSites(w, r)
		default:
			fmt.Fprintf(w, "GET METHOD: %s DOES NOT EXIST", path)
		}

	case "POST":
		switch path {
		case "signup":
			signup(w, r)
		case "login":
			login(w, r)

		default:
			fmt.Fprintf(w, "POST METHOD: %s DOES NOT EXIST", path)
		}

	default:
		fmt.Fprintf(w, "METHOD: %s IS NOT ALLOWED", method)
	}
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:5000")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")

}

func main() {
	http.HandleFunc("/", router)
	http.ListenAndServe(":8080", nil)
}
