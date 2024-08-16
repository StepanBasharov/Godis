package api

import (
	"github.com/gin-gonic/gin"
	"godis/internal/connections/http/handlers"
	"godis/internal/storage"
)

func SetupApi(engine *gin.Engine, storage *storage.Storage) {
	// storage api
	storageApi := engine.Group("/storage")
	storageApi.POST("/set", func(context *gin.Context) {
		handlers.StoreHandlerSet(context, storage)
	})
	storageApi.GET("/get", func(context *gin.Context) {
		handlers.StoreHandlerGet(context, storage)
	})
}
