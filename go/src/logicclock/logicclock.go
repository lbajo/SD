package logicclock

import "fmt"
import "encoding/json"
import "errors"

type Clock struct {
	Name string
	Log  string
	Map  map[string]int
}

func NewClock(Name string) Clock {
	mapClocks := make(map[string]int)
	mapClocks[Name] = 0
	return Clock{Name, "", mapClocks}
}

func AddClock(clock Clock) {
	_, ok := clock.Map[clock.Name]

	if ok {
		clock.Map[clock.Name]++
	} else {
		clock.Map[clock.Name] = 1
	}
}

func GetClockValue(clock Clock, name string) int {
	return clock.Map[name]
}

func InsertClock(clock Clock, name string, value int) {
	clock.Map[name] = value
}

/*
func SetClock(c1 Clock, c2 Clock) {
	for key, val1 := range c1.Map {
		val2 := c2.Map[key]
		if val2 > val1 {
			c1.Map[key] = val2
		}
	}
	for key, val2 := range c2.Map {
		val1 := c1.Map[key]
		if val2 > val1 {
			c1.Map[key] = val2
		}
	}
}*/

func SetClock(c1 Clock, c2 Clock) {
	for key, val1 := range c1.Map {
		val2 := GetClockValue(c2, key)
		if val2 > val1 {
			InsertClock(c1, key, val2)
		}
	}
	for key, val2 := range c2.Map {
		val1 := GetClockValue(c1, key)
		if val2 > val1 {
			InsertClock(c1, key, val2)
		}
	}
}

func Max(c1 Clock, c2 Clock) {
	SetClock(c1, c2)
	SetClock(c2, c1)
}

func DeleteClock(clock Clock) {
	_, ok := clock.Map[clock.Name]
	if ok {
		delete(clock.Map, clock.Name)
	}
}

func ToJSON(clock Clock) string {
	value, err := json.Marshal(clock)
	if err != nil {
		fmt.Println(errors.New("Error al convertir a JSON"))
		return ""
	}
	return string(value)
}

func FromJSON(txt string) Clock {
	var clock Clock
	err := json.Unmarshal([]byte(txt), &clock)
	if err != nil {
		fmt.Println(errors.New("Error al convertir a string"))
	}
	return clock
}
