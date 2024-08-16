package main

import (
	"godis/internal/connections/http"
	"godis/internal/storage"
	"log/slog"
	"strconv"
	"sync"
)

func InitHttpServer(wg *sync.WaitGroup, systemStorage *storage.Storage, countInstance int) {
	// init http server connections
	for i := 0; i < countInstance; i++ {
		httpServer := http.NewHttpServer(systemStorage)
		wg.Add(1)
		port := ":" + strconv.Itoa(8340+i)
		go httpServer.StartHttpServer(wg, port)
		slog.Info("INIT http server on port %s", port)
	}

}

func main() {
	// wait group for many instance
	var wg sync.WaitGroup
	// sync.Map for stored data
	var systemMap sync.Map
	systemStorage := storage.NewStorage(&systemMap)
	// init http server
	InitHttpServer(&wg, &systemStorage, 10)
	wg.Wait()
}
