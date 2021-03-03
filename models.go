package main

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

var dbUsers = map[string]user{}
var dbSites = []kvSite{}
var dbChannels = []kvChannel{}
var dbPlaylists = []kvPlaylist{}
var dbPlaylistChannelLinks = []kvPlaylistChannelLink{}
var dbTracks = []kvTrack{}

type user struct {
	ID       int
	Name     string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type session struct {
	un         string
	lastActive time.Time
}

type kvSite struct {
	ID       int         `json:"id"`
	UserID   int         `json:"user_id"`
	Name     string      `json:"name"`
	Channels []kvChannel `json:"channels"`
}

type kvChannel struct {
	ID        int          `json:"id"`
	SiteID    int          `json:"site_id"`
	Name      string       `json:"name"`
	Playlists []kvPlaylist `json:"playlists"`
}

type kvPlaylist struct {
	ID     int       `json:"id"`
	Name   string    `json:"name"`
	Tracks []kvTrack `json:"tracks"`
}

type kvPlaylistChannelLink struct {
	PlaylistID int
	ChannelID  int
}

type kvTrack struct {
	ID        int    `json:"id"`
	PlayistID int    `json:"playlist_id"`
	Name      string `json:"name"`
}

func initDB() {
	makeDBUsers()
	makeDBSites()
	makeDBChannels()
	makeDBPlaylistChannelLinks()
	makeDBPlaylists()
	makeBDTracks()
}

func makeDBUsers() {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}
	dbUsers["avi"] = user{1, "avi", "avi@gmai.com", string(hashedPassword)}
	dbUsers["tsachi"] = user{2, "tsachi", "tsachi@gmai.com", string(hashedPassword)}
	dbUsers["cheli"] = user{3, "cheli", "cheli@gmai.com", string(hashedPassword)}
	dbUsers["rutie"] = user{4, "rutie", "rutie@gmai.com", string(hashedPassword)}
	dbUsers["tara"] = user{5, "tara", "tara@gmai.com", string(hashedPassword)}
}

func makeDBSites() {
	dbSites = append(dbSites,
		kvSite{1, 1, "site:1", []kvChannel{}},
		kvSite{2, 1, "site:2", []kvChannel{}},
		kvSite{3, 3, "site:3", []kvChannel{}},
		kvSite{4, 1, "site:4", []kvChannel{}})
}

func makeDBChannels() {
	dbChannels = append(dbChannels,
		kvChannel{1, 1, "channel:1,site:1", []kvPlaylist{}},
		kvChannel{2, 1, "channel:2:site1", []kvPlaylist{}},
		kvChannel{3, 2, "channel:1,site:2", []kvPlaylist{}},
		kvChannel{4, 2, "channel:2:site2", []kvPlaylist{}})

}

func makeDBPlaylists() {
	dbPlaylists = append(dbPlaylists,
		kvPlaylist{1, "playlist:1", []kvTrack{}},
		kvPlaylist{2, "playlist:2", []kvTrack{}},
		kvPlaylist{3, "playlist:3", []kvTrack{}})
}

func makeDBPlaylistChannelLinks() {
	dbPlaylistChannelLinks = append(dbPlaylistChannelLinks,
		kvPlaylistChannelLink{1, 1},
		kvPlaylistChannelLink{2, 1},
		kvPlaylistChannelLink{3, 2},
		kvPlaylistChannelLink{1, 2})
}

func makeBDTracks() {
	dbTracks = append(dbTracks,
		kvTrack{1, 1, "track1:playlist1"},
		kvTrack{2, 1, "track2:playlist1"},
		kvTrack{3, 1, "track3:playlist1"},
		kvTrack{4, 2, "track1:playlist2"},
		kvTrack{5, 2, "track2:playlist2"},
		kvTrack{6, 2, "track3:playlist2"},
		kvTrack{7, 3, "track1:playlist3"})

}
