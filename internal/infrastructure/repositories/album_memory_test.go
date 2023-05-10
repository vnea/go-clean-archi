package repositories

import (
	"github.com/vnea/go-clean-archi/internal/domain/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAll_ShouldReturnAllStoredAlbums(t *testing.T) {
	// Given
	storedAlbums := []AlbumMemory{
		{
			ID:   "1",
			Name: "You rock!",
		},
	}
	repository := AlbumMemoryRepository{
		Albums: storedAlbums,
	}

	// When
	returnedAlbums := repository.GetAll()

	// Then
	expectedAlbums := []entities.Album{
		{
			ID:   "1",
			Name: "You rock!",
		},
	}
	assert.Equal(t, expectedAlbums, returnedAlbums)
}
