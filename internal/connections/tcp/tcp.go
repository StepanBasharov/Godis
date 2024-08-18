package tcp

import (
	"godis/internal/connections/tcp/handlers"
	"godis/internal/storage"
	"log/slog"
	"net"
	"sync"
)

type ServerTCP struct {
	storageMap *storage.Storage
	port       string
}

func NewTcpServer(s *storage.Storage, port string) ServerTCP {
	return ServerTCP{storageMap: s, port: port}
}

func (s *ServerTCP) StartHttpServer(wg *sync.WaitGroup) {
	defer wg.Done()
	slog.Info("Tcp server started", "port", s.port)
	server, err := net.Listen("tcp", s.port)
	if err != nil {
		panic(err)
	}
	defer server.Close()
	for {
		conn, err := server.Accept()
		if err != nil {
			panic(err)
		}
		go handlers.HandleConnection(conn, s.storageMap)
	}
}
