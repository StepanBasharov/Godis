package http

import (
	"github.com/gin-gonic/gin"
	"godis/internal/connections/http/api"
	"godis/internal/storage"
	"sync"
)

type ServerHTTP struct {
	engine *gin.Engine
	port   string
}

func NewHttpServer(s *storage.Storage, port string) ServerHTTP {
	engine := gin.Default()
	api.SetupApi(engine, s)

	return ServerHTTP{engine, port}
}

func (s *ServerHTTP) StartServer(wg *sync.WaitGroup) {
	defer wg.Done()
	gin.SetMode(gin.ReleaseMode)
	s.engine.Run(s.port)
}
