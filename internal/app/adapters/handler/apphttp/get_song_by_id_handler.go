package apphttp

import (
	"back-vozes-cifras/internal/app/application/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// GetSongByIDHandler returns a song by id
func (h *Handler) GetSongByIDHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("ID is required"))
			return
		}
		idInt, err := strconv.Atoi(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("ID must be a number"))
			return
		}
		song, err := h.useCase.GetSongByIDUseCase(idInt)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintln("Error: ", err)))
			return
		}
		response := mapSongToResponse(song)
		responseJSON, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(responseJSON)
	})
}

func mapSongToResponse(song *model.Song) *getSongByIDResponse {
	return &getSongByIDResponse{
		ID:     song.ID,
		Title:  song.Title,
		Artist: song.Artist,
		Tags:   song.Tags,
		Verses: mapVersesToResponse(song.Verses),
	}
}

func mapVersesToResponse(verses []model.Verse) []verse {
	var response []verse
	for _, v := range verses {
		response = append(response, verse{
			Number: v.Number,
			Chorus: v.Chorus,
			Verse:  v.Verse,
			Chords: mapChordsToResponse(v.Chords),
		})
	}
	return response
}

func mapChordsToResponse(chords []model.Chord) []chord {
	var response []chord
	for _, c := range chords {
		response = append(response, chord{
			Number:   c.Number,
			Chord:    c.Chord,
			Position: c.Position,
		})
	}
	return response
}

type getSongByIDResponse struct {
	ID     int      `json:"id"`
	Title  string   `json:"title"`
	Artist string   `json:"artist"`
	Tags   []string `json:"tags"`
	Verses []verse  `json:"verses"`
}

type verse struct {
	Number int     `json:"number"`
	Chorus bool    `json:"chorus"`
	Verse  string  `json:"verse"`
	Chords []chord `json:"chords"`
}

type chord struct {
	Number   int    `json:"number"`
	Chord    string `json:"chord"`
	Position int    `json:"position"`
}
