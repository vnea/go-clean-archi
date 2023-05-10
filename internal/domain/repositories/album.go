package repositories

import "github.com/vnea/go-clean-archi/internal/domain/entities"

type AlbumRepository interface {
	GetAll() []entities.Album
}
