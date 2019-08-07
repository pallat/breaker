package main

import (
	"fmt"

	"github.com/pallat/breaker/queue"
	"github.com/pallat/breaker/server"
)

func main() {
	q := queue.New()

	go server.Listen(q)

	for s := range q.Success() {
		fmt.Println(s)
	}
}
