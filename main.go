package main

import (
	"github.com/vnea/go-clean-archi/internal/application/controllers"
	"github.com/vnea/go-clean-archi/internal/domain/usecases"
	"github.com/vnea/go-clean-archi/internal/infrastructure/repositories"
)

func main() {
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
	var albumController = controllers.AlbumController{
		GetAllAlbumUseCase: getAllAlbumsUseCase,
	}

	router := controllers.SetupRouter(albumController)

	router.Run("localhost:8080")
}
