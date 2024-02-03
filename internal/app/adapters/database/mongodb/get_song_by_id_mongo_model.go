package mongodb

import "back-vozes-cifras/internal/app/application/model"

func (s *song) toModel() *model.Song {
	verses := make([]model.Verse, 0)
	for _, v := range s.Verses {
		verses = append(verses, v.toModel())
	}

	return &model.Song{
		ID:     s.ID,
		Title:  s.Title,
		Artist: s.Artist,
		Tags:   s.Tags,
		Verses: verses,
	}
}

func (v *verse) toModel() model.Verse {
	chords := make([]model.Chord, 0)
	for _, c := range v.Chords {
		chords = append(chords, c.toModel())
	}

	return model.Verse{
		Number: v.Number,
		Chorus: v.Chorus,
		Verse:  v.Verse,
		Chords: chords,
	}
}

func (c *chord) toModel() model.Chord {
	return model.Chord{
		Number:   c.Number,
		Chord:    c.Chord,
		Position: c.Position,
	}
}

type song struct {
	ID     int      `bson:"_id"`
	Title  string   `bson:"title"`
	Artist string   `bson:"artist"`
	Tags   []string `bson:"tags"`
	Verses []verse  `bson:"verses"`
}

type verse struct {
	Number int     `bson:"number"`
	Chorus bool    `bson:"chorus"`
	Verse  string  `bson:"verse"`
	Chords []chord `bson:"chords"`
}

type chord struct {
	Number   int    `bson:"number"`
	Chord    string `bson:"chord"`
	Position int    `bson:"position"`
}
