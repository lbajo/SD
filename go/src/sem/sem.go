//LORENA BAJO REBOLLO		TELEMÁTICA

package sem

import (
	"errors"
	"fmt"
	"sync"
)

type UpDowner interface {
	Up()
	Down()
}

type Sem struct {
	c *sync.Cond
	n int
}

var mutex sync.Mutex

func NewSem(ntok int) *Sem {
	if ntok < 0 {
		fmt.Println(errors.New("Imposible crear con un número menor que 0"))
		return nil
	}

	c := sync.NewCond(&mutex)
	sem := Sem{c, ntok}

	return &sem
}

func (s *Sem) Up() {
	s.c.L.Lock()
	s.n++
	s.c.Signal()
	s.c.L.Unlock()
}

func (s *Sem) Down() {
	s.c.L.Lock()
	if s.n <= 0 {
		s.c.Wait()
	}
	s.n--
	s.c.L.Unlock()
}
