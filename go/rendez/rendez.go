//LORENA BAJO REBOLLO		TELEMÁTICA

package rendez

import (
	"sync"
)

type Info struct {
	wg  sync.WaitGroup
	val interface{}
}

var m = make(map[int]*Info)
var mutex sync.Mutex

func Rendezvous(tag int, val interface{}) interface{} {

	mutex.Lock()

	elem, ok := m[tag]

	if ok {
		val, elem.val = elem.val, val
		m[tag].wg.Done()
		delete(m, tag)
		mutex.Unlock()
		return val

	} else {
		if elem == nil {
			elem = new(Info)
		}

		elem.wg.Add(1)
		elem.val = val
		m[tag] = elem
		mutex.Unlock()
		elem.wg.Wait()

		return elem.val
	}
}
