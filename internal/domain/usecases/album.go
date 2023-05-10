package usecases

import (
	"github.com/vnea/go-clean-archi/internal/domain/entities"
	"github.com/vnea/go-clean-archi/internal/domain/repositories"
)

type GetAllAlbums struct {
	AlbumRepository repositories.AlbumRepository
}

func (useCase GetAllAlbums) Execute() []entities.Album {
	return useCase.AlbumRepository.GetAll()
}
