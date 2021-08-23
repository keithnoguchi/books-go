package ch07

import (
	"fmt"
	"io"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

type ByArtist []*Track

func (b ByArtist) Len() int { return len(b) }
func (b ByArtist) Less(i, j int) bool { return b[i].Artist < b[j].Artist }
func (b ByArtist) Swap(i, j int) { b[i], b[j] = b[j], b[i] }

type ByCustom struct {
	Tracks   []*Track
	LessFunc func(i, j *Track) bool
}

func (b ByCustom) Len() int { return len(b.Tracks) }
func (b ByCustom) Less(i, j int) bool { return b.LessFunc(b.Tracks[i], b.Tracks[j]) }
func (b ByCustom) Swap(i, j int) { b.Tracks[i], b.Tracks[j] = b.Tracks[j], b.Tracks[i] }

func PrintTracks(out io.Writer, tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	var tw = new(tabwriter.Writer).Init(out, 0, 8, 2, ' ', 0)

	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()
}
