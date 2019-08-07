package server

import (
	"bytes"
	"io"
	"log"
	"net"

	"github.com/pallat/breaker/queue"
	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("port", "12001")

	viper.AutomaticEnv()
}

func Listen(q *queue.Queue, pause chan struct{}) <-chan net.Conn {
	l, err := net.Listen("tcp", ":"+viper.GetString("port"))
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		// Wait for a connection.
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.
		go func(c net.Conn) {
			// Echo all incoming data.
			buf := new(bytes.Buffer)
			io.Copy(buf, c)
			// Shut down the connection.

			q.Input() <- buf.String()

			c.Close()
		}(conn)
	}
}
