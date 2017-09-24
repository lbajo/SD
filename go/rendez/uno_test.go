package rendez

import (
	"testing"
	"fmt"
	"time"
)

func lanzar1(tag int,val interface{}){
	r1 := Rendezvous(tag,val)
	fmt.Println("Thread: 1 Tag:",tag,"-->",r1)
}

func lanzar2(tag int,val interface{}){
	r1 := Rendezvous(tag,val)
	fmt.Println("Thread: 2 Tag:",tag,"-->",r1)
}

func TestUno(t* testing.T){
	go lanzar1(1,"hola")
	time.Sleep(1000*time.Millisecond)
	go lanzar2(1,"caracola")
	time.Sleep(2000*time.Millisecond)
}