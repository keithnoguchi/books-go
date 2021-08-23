package main

import (
	"os"
	"sort"
	"time"

	"book/ch07"
)

var tracks = []*ch07.Track {
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func main() {
	ch07.PrintTracks(os.Stdout, tracks)
	tracksByArtist := ch07.ByArtist(tracks)
	sort.Sort(tracksByArtist)
	ch07.PrintTracks(os.Stdout, tracksByArtist)
	tracksByCustom := ch07.ByCustom{
		tracks,
		func(i, j *ch07.Track) bool {
			if i.Title < j.Title {
				return true
			}
			if i.Artist < j.Artist {
				return true
			}
			if i.Album < j.Album {
				return true
			}
			return false
		},
	}
	sort.Sort(tracksByCustom)
	ch07.PrintTracks(os.Stdout, tracksByCustom.Tracks)
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(err)
	}
	return d
}
