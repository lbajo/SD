package rendez

import (
	"testing"
	"fmt"
	"time"
)

func lanzardisttag(i int, tag int,val interface{}){
	r1 := Rendezvous(tag,val)
	fmt.Println("Thread",i,"Tag:",tag,"-->",r1)
}

func TestCuatro(t* testing.T){

	for i:=0; i<4; i++{
		go lanzardisttag(i,4,i)
	}
	for i:=4; i<8; i++{
		go lanzardisttag(i,5,i)
	}
	time.Sleep(5000*time.Millisecond)
}