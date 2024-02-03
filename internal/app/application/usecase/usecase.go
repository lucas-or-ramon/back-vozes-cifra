package usecase

import (
	"back-vozes-cifras/internal/app/application/model"
	"back-vozes-cifras/internal/app/application/repository"
)

type UseCase interface {
	// GetSongByIDUseCase returns a song by id
	GetSongByIDUseCase(id int) (*model.Song, error)
}

type useCase struct {
	repo repository.Repository
}

func NewUseCase(repo repository.Repository) UseCase {
	return &useCase{
		repo: repo,
	}
}
