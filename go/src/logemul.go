package main

import (
	"logiclog"
	"os"
	"time"
)

type Friend struct {
	ch   chan logiclog.Msg
	file *os.File
	name string
}

var (
	a1 = Friend{make(chan logiclog.Msg), createFile("Amigo1"), "Amigo1"}
	a2 = Friend{make(chan logiclog.Msg), createFile("Amigo2"), "Amigo2"}
	a3 = Friend{make(chan logiclog.Msg), createFile("Amigo3"), "Amigo3"}
	a4 = Friend{make(chan logiclog.Msg), createFile("Amigo4"), "Amigo4"}
)

func createFile(name string) *os.File {
	f, err := os.Create("/tmp/" + name)
	if err != nil {
		panic(err)
	}
	return f
}

func friend1() {
	log1 := logiclog.NewLog(a1.name, a1.file)
	logiclog.WriteLog(log1, "Tengo ganas de salir")

	msg := logiclog.SendLog(log1, "Sales hoy 2?")
	a2.ch <- msg

	msg_rec := <-a1.ch
	logiclog.ReceiveLog(log1, msg_rec)

	msg = logiclog.SendLog(log1, "OK1")
	a2.ch <- msg

	msg_rec = <-a1.ch
	logiclog.ReceiveLog(log1, msg_rec)

	msg = logiclog.SendLog(log1, "2, quedamos en la puerta")
	a2.ch <- msg

	msg = logiclog.SendLog(log1, "3, quedamos en la puerta")
	a3.ch <- msg

	msg_rec = <-a1.ch
	logiclog.ReceiveLog(log1, msg_rec)

	msg_rec = <-a1.ch
	logiclog.ReceiveLog(log1, msg_rec)

}

func friend2() {
	log2 := logiclog.NewLog(a2.name, a2.file)
	logiclog.WriteLog(log2, "Menuda tarde de lluvia")

	msg_rec := <-a2.ch
	logiclog.ReceiveLog(log2, msg_rec)

	msg := logiclog.SendLog(log2, "Vale, se lo voy a preguntar a 3")
	a1.ch <- msg

	msg_rec = <-a2.ch
	logiclog.ReceiveLog(log2, msg_rec)

	msg = logiclog.SendLog(log2, "Sales hoy 3?")
	a3.ch <- msg

	msg_rec = <-a2.ch
	logiclog.ReceiveLog(log2, msg_rec)

	msg = logiclog.SendLog(log2, "OK2")
	a3.ch <- msg

	msg_rec = <-a2.ch
	logiclog.ReceiveLog(log2, msg_rec)

	msg = logiclog.SendLog(log2, "OK21")
	a3.ch <- msg

	msg = logiclog.SendLog(log2, "3 sale, 4 no")
	a1.ch <- msg

	msg_rec = <-a2.ch
	logiclog.ReceiveLog(log2, msg_rec)

	msg = logiclog.SendLog(log2, "Vale, soy 2")
	a1.ch <- msg

}
func friend3() {
	log3 := logiclog.NewLog(a3.name, a3.file)
	logiclog.WriteLog(log3, "Me aburro")

	msg_rec := <-a3.ch
	logiclog.ReceiveLog(log3, msg_rec)

	msg := logiclog.SendLog(log3, "Vale, se lo voy a preguntar a 4")
	a2.ch <- msg

	msg_rec = <-a3.ch
	logiclog.ReceiveLog(log3, msg_rec)

	msg = logiclog.SendLog(log3, "Sales hoy 4?")
	a4.ch <- msg

	msg_rec = <-a3.ch
	logiclog.ReceiveLog(log3, msg_rec)

	msg = logiclog.SendLog(log3, "OK3")
	a4.ch <- msg

	msg = logiclog.SendLog(log3, "4 no viene, cine?")
	a2.ch <- msg

	msg_rec = <-a3.ch
	logiclog.ReceiveLog(log3, msg_rec)

	msg_rec = <-a3.ch
	logiclog.ReceiveLog(log3, msg_rec)

	msg = logiclog.SendLog(log3, "Vale, soy 3")
	a1.ch <- msg

}

func friend4() {
	log4 := logiclog.NewLog(a4.name, a4.file)
	logiclog.WriteLog(log4, "Tengo sueño")

	msg_rec := <-a4.ch
	logiclog.ReceiveLog(log4, msg_rec)

	msg := logiclog.SendLog(log4, "Yo no salgo")
	a3.ch <- msg

	msg_rec = <-a4.ch
	logiclog.ReceiveLog(log4, msg_rec)

	logiclog.WriteLog(log4, "A dormir")
}

func main() {

	go friend1()
	time.Sleep(50 * time.Millisecond)
	go friend2()
	time.Sleep(50 * time.Millisecond)
	go friend3()
	time.Sleep(50 * time.Millisecond)
	go friend4()

	time.Sleep(10000 * time.Millisecond)
}
