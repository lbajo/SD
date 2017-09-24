package rendez

import (
	"testing"
	"fmt"
	"time"
)

func lanzariguales(i int, tag int,val interface{}){
	r1 := Rendezvous(tag,val)
	fmt.Println("Thread",i,"Tag:",tag,"-->",r1)
}

func TestDos(t* testing.T){

	for i:=0; i<5; i++{
		go lanzariguales(i,2,"dos")
	}

	time.Sleep(3000*time.Millisecond)
}