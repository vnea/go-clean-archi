package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/vnea/go-clean-archi/internal/domain/usecases"
	"github.com/vnea/go-clean-archi/internal/infrastructure/repositories"
)

var albumMemories = []repositories.AlbumMemory{
	{
		ID:   "1",
		Name: "You rock!",
	},
}

var albumRepository = repositories.AlbumMemoryRepository{
	Albums: albumMemories,
}
var getAllAlbumsUseCase = usecases.GetAllAlbums{
	AlbumRepository: albumRepository,
}
var albumController = AlbumController{
	GetAllAlbumUseCase: getAllAlbumsUseCase,
}

func TestRouteGetAlbums_ShouldReturnAlbums(t *testing.T) {
	// Given
	router := SetupRouter(albumController)

	responseWriter := httptest.NewRecorder()
	request, _ := http.NewRequest(
		"GET",
		"/albums",
		nil,
	)

	// When
	router.ServeHTTP(responseWriter, request)

	// Then
	assert.Equal(t, http.StatusOK, responseWriter.Code)

	var httpResponse []AlbumApi
	_ = json.Unmarshal(responseWriter.Body.Bytes(), &httpResponse)
	expectedHttpResponse := []AlbumApi{
		{
			ID:   "1",
			Name: "You rock!",
		},
	}
	assert.Equal(t, expectedHttpResponse, httpResponse)
}
