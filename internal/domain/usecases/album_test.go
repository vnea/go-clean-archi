package usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/vnea/go-clean-archi/internal/domain/entities"
)

var expectedAlbums = []entities.Album{
	{
		ID:   "1",
		Name: "You rock!",
	},
}

type AlbumRepositoryMock struct{}

func (m AlbumRepositoryMock) GetAll() []entities.Album {
	return expectedAlbums
}

func TestExecute_ShouldReturnTheAlbums(t *testing.T) {
	// Given
	albumRepositoryMock := new(AlbumRepositoryMock)
	getAllAlbums := GetAllAlbums{
		AlbumRepository: albumRepositoryMock,
	}

	// When
	returnedAlbums := getAllAlbums.Execute()

	// Then
	assert.Equal(t, expectedAlbums, returnedAlbums)
}
