package repository

import "back-vozes-cifras/internal/app/application/model"

type Repository interface {
	// GetSongByIDRepository returns a song by id
	GetSongByIDRepository(id int) (*model.Song, error)
}
