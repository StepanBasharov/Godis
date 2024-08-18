package handlers

import (
	"github.com/gin-gonic/gin"
	"godis/internal/connections/http/data"
	"godis/internal/storage"
	"net/http"
)

func StoreHandlerSet(ctx *gin.Context, storage *storage.Storage) {
	// handler for stored data {key : val }
	var content data.AddToStore
	if err := ctx.ShouldBindJSON(&content); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	err := storage.Set(content.Key, content.Val)
	if err != nil {
		ctx.String(http.StatusInternalServerError, "ERROR")
	}

	ctx.String(http.StatusOK, "OK")
}

func StoreHandlerGet(ctx *gin.Context, storage *storage.Storage) {
	// handler for getting data
	key := ctx.Query("key")
	value := storage.Get(key)
	ctx.JSON(http.StatusOK, data.GetFromStore{Val: value})
}
