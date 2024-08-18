package connections

import (
	"sync"
)

type Server interface {
	StartHttpServer(wg *sync.WaitGroup)
}
