package sem

import (
	"fmt"
	"testing"
	"time"
)

const N = 2

var s = NewSem(N)

var i int = 0

func CogerFicha() {
	s.Down()
	i++
}

func TestDos(t *testing.T) {
	go CogerFicha()
	go CogerFicha()
	go CogerFicha()

	time.Sleep(1000 * time.Millisecond)

	if i != 2 {
		t.Error()
	}
	fmt.Println("i=2, continuamos")

	s.Up()

	time.Sleep(1000 * time.Millisecond)

	if i != 3 {
		t.Error()
	}

	fmt.Println("i=3, todo en orden, fin")

}
