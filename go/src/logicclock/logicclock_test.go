package logicclock

import (
    "testing"
    "fmt"
)

func TestLogicclock(t *testing.T){
	clock1 := NewClock("Prueba1")
    clock2 := NewClock("Prueba2")

    AddClock(clock1)
    AddClock(clock1)
 
    if GetClockValue(clock1) != 2{
        t.Error()
    }

    InsertClock(clock2,2)

    if GetClockValue(clock2) != 2{
        t.Error()
    }

    fmt.Println("Clock1 (toJson)",ToJSON(clock1))
    json := ToJSON(clock2)
    fmt.Println("Clock2 (fromJson)",FromJSON(json))

}