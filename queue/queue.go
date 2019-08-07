package queue

import "fmt"

type Queue struct {
	msg chan string
	ok  chan string
	err chan string
}

func New() *Queue {
	q := &Queue{
		msg: make(chan string),
		ok:  make(chan string),
	}
	go Send(q.msg, q.ok)
	return q
}

func (q *Queue) Input() chan<- string {
	return q.msg
}

func (q *Queue) Success() <-chan string {
	return q.ok
}

func (q *Queue) Errors() <-chan string {
	return q.err
}

func Send(msg chan string, ok chan string) {
	for {
		fmt.Println(<-msg)
		ok <- "done"
	}
}
