package usecase

import (
	"back-vozes-cifras/internal/app/application/model"
	"errors"
)

func (u *useCase) GetSongByIDUseCase(id int) (*model.Song, error) {
	if id == 0 {
		return nil, errors.New("id is required")
	}

	return u.repo.GetSongByIDRepository(id)
}
