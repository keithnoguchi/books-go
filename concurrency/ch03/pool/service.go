// sync.Pool example, page 61
package pool

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

func StartColdNetworkDaemon(proto, addr string) *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		server, err := net.Listen(proto, addr)
		if err != nil {
			log.Fatalf("cannot listen: %v", err)
		}
		defer server.Close()
		wg.Done()
		for {
			conn, err := server.Accept()
			if err != nil {
				log.Printf("cannot accept connection: %v", err)
				continue
			}
			connectToService()
			fmt.Fprintln(conn, "")
			conn.Close()
		}
	}()
	return &wg
}

func StartWarmNetworkDaemon(proto, addr string) *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		connPool := warmServiceConnPool(10)
		server, err := net.Listen(proto, addr)
		if err != nil {
			log.Fatalf("cannot listen: %v", err)
		}
		defer server.Close()
		wg.Done()
		for {
			conn, err := server.Accept()
			if err != nil {
				log.Printf("cannot accept connection: %v", err)
				continue
			}
			svcConn := connPool.Get()
			fmt.Fprintln(conn, "")
			connPool.Put(svcConn)
			conn.Close()
		}
	}()
	return &wg
}

// connection pool.
func warmServiceConnPool(num int) *sync.Pool {
	p := &sync.Pool{
		New: connectToService,
	}
	for i := 0; i < num; i++ {
		p.Put(p.New())
	}
	return p
}

// slow service simulation.
func connectToService() interface{} {
	time.Sleep(1 * time.Second)
	return struct{}{}
}
