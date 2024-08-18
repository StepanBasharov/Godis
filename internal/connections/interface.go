package connections

import (
	"sync"
)

type Server interface {
	StartServer(wg *sync.WaitGroup)
}
