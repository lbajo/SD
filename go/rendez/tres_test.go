package rendez

import (
	"testing"
	"fmt"
	"time"
)

func lanzardistintos(i int, tag int,val interface{}){
	r1 := Rendezvous(tag,val)
	fmt.Println("Thread",i,"Tag:",tag,"-->",r1)
}

func TestTres(t* testing.T){

	for i:=0; i<5; i++{
		go lanzardistintos(i,3,i)
	}

	time.Sleep(3000*time.Millisecond)
}