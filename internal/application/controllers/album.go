package controllers

import (
	"github.com/vnea/go-clean-archi/internal/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/vnea/go-clean-archi/internal/domain/usecases"
)

type AlbumApi struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type AlbumController struct {
	GetAllAlbumUseCase usecases.GetAllAlbums
}

func (controller AlbumController) GetAlbums(c *gin.Context) {
	albums := controller.GetAllAlbumUseCase.Execute()
	albumApis := controller.toAlbumApis(albums)

	c.JSON(http.StatusOK, albumApis)
}

func (controller AlbumController) toAlbumApis(albums []entities.Album) []AlbumApi {
	var albumApis []AlbumApi

	for _, album := range albums {
		albumApis = append(albumApis, AlbumApi{
			ID:   album.ID,
			Name: album.Name,
		})
	}

	return albumApis
}
