// sync.Pool example, page 61
package pool

import (
	"io/ioutil"
	"net"
	"testing"
)

const (
	proto      = "tcp"
	coldDaemon = "localhost:8080"
	warmDaemon = "localhost:8081"
)

func init() {
	coldDaemonStarted := StartColdNetworkDaemon(proto, coldDaemon)
	warmDaemonStarted := StartWarmNetworkDaemon(proto, warmDaemon)
	coldDaemonStarted.Wait()
	warmDaemonStarted.Wait()
}

func BenchmarkColdNetworkRequest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		conn, err := net.Dial(proto, coldDaemon)
		if err != nil {
			b.Fatalf("cannot dial host: %v", err)
		}
		if _, err := ioutil.ReadAll(conn); err != nil {
			b.Fatalf("cannot read: %v", err)
		}
		conn.Close()
	}
}

func BenchmarkWarmNetworkRequest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		conn, err := net.Dial(proto, warmDaemon)
		if err != nil {
			b.Fatalf("cannot dial host: %v", err)
		}
		if _, err := ioutil.ReadAll(conn); err != nil {
			b.Fatalf("cannot read: %v", err)
		}
		conn.Close()
	}
}
