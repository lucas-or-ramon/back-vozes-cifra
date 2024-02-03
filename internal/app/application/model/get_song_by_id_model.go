package model

type Song struct {
	ID     int
	Title  string
	Artist string
	Tags   []string
	Verses []Verse
}

type Verse struct {
	Number int
	Chorus bool
	Verse  string
	Chords []Chord
}

type Chord struct {
	Number   int
	Chord    string
	Position int
}
