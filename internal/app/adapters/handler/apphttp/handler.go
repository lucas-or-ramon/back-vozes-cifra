package apphttp

import "back-vozes-cifras/internal/app/application/usecase"

type Handler struct {
	useCase usecase.UseCase
}

// NewHandler creates a new handler
func NewHandler(useCase usecase.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}
