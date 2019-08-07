package switching

import (
	"fmt"
	"testing"
)

func TestX(t *testing.T) {
	reset, red, blue := Switch()

	reset()

	fmt.Println("red")
	red() <- struct{}{}

	fmt.Println("blue")
	<-blue()

	reset()

	fmt.Println("blue")
	blue() <- struct{}{}

	fmt.Println("red")
	<-red()

	fmt.Println("end")

}
