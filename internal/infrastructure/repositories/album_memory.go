package repositories

import "github.com/vnea/go-clean-archi/internal/domain/entities"

type AlbumMemory struct {
	ID   string
	Name string
}

type AlbumMemoryRepository struct {
	Albums []AlbumMemory
}

func (repository AlbumMemoryRepository) GetAll() []entities.Album {
	var albums []entities.Album

	for _, albumMemory := range repository.Albums {
		album := entities.Album{
			ID:   albumMemory.ID,
			Name: albumMemory.Name,
		}
		albums = append(albums, album)
	}

	return albums
}
