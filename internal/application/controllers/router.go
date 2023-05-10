package controllers

import "github.com/gin-gonic/gin"

func setupRouter(albumController AlbumController) *gin.Engine {
	router := gin.Default()

	router.GET("/albums", albumController.GetAlbums)

	return router
}
