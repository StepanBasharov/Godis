package main

import (
	"godis/internal/connections/http"
	"godis/internal/connections/tcp"
	"godis/internal/storage"
	"log/slog"
	"strconv"
	"sync"
)

func initHttpServer(wg *sync.WaitGroup, systemStorage *storage.Storage, countInstance int) {
	// init http server connections
	for i := 0; i < countInstance; i++ {
		port := ":" + strconv.Itoa(8340+i)
		httpServer := http.NewHttpServer(systemStorage, port)
		wg.Add(1)
		go httpServer.StartServer(wg)
		slog.Info("INIT http server on port", "port", port)
	}

}

func initTcpServer(wg *sync.WaitGroup, systemStorage *storage.Storage, countInstance int) {
	for i := 0; i < countInstance; i++ {
		port := ":" + strconv.Itoa(32410+i)
		tcpServer := tcp.NewTcpServer(systemStorage, port)
		wg.Add(1)
		go tcpServer.StartServer(wg)
		slog.Info("INIT tcp server on port", "port", port)
	}
}

func main() {
	// wait group for many instance
	var wg sync.WaitGroup
	// sync.Map for stored data
	var systemMap sync.Map
	systemStorage := storage.NewStorage(&systemMap)
	// init http server
	initHttpServer(&wg, &systemStorage, 10)
	initTcpServer(&wg, &systemStorage, 10)
	wg.Wait()

}
