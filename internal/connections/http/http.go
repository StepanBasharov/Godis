package http

import (
	"github.com/gin-gonic/gin"
	"godis/internal/connections/http/api"
	"godis/internal/storage"
	"sync"
)

type Server struct {
	engine  *gin.Engine
	storage *storage.Storage
}

func NewHttpServer(s *storage.Storage) Server {
	engine := gin.Default()
	api.SetupApi(engine, s)

	return Server{engine, s}
}

func (s *Server) StartHttpServer(wg *sync.WaitGroup, port string) {
	defer wg.Done()
	gin.SetMode(gin.ReleaseMode)
	s.engine.Run(port)
}
