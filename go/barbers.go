package main

import (
	"fmt"
	"strconv"
	"time"
)

const (
	MAXBARBEROS          = 2
	MAXCLIENTESESPERANDO = 5
)

type Cliente struct {
	ch chan string
	id int
}

type Barbero struct {
	id int
	ch chan Cliente
}

type Recepcion struct {
	recepcionista chan Cliente
	sala_espera   chan string
}

var (
	clientesesperando int = 0
	barberosocup      int = 0

	b1    = Barbero{1, make(chan Cliente)}
	b2    = Barbero{2, make(chan Cliente)}
	recep = Recepcion{make(chan Cliente), make(chan string)}
)

func barbero(barb Barbero) {
	msg1_bar := "Barbero " + strconv.Itoa(barb.id) + ": me duermo esperando clientes\n"
	msg2_bar := "Barbero " + strconv.Itoa(barb.id) + ": empiezo a cortar el pelo\n"
	msg3_bar := "Barbero " + strconv.Itoa(barb.id) + ": termino de cortar el pelo\n"

	for {
		fmt.Printf(msg1_bar)
		cl := <-barb.ch
		barberosocup++
		fmt.Printf(msg2_bar)
		cl.ch <- "cortando"
		time.Sleep(500 * time.Millisecond)
		fmt.Printf(msg3_bar)
		cl.ch <- "termino"
		barberosocup--

	}
}

func cliente(ncli int) {

	cl := Cliente{make(chan string), ncli}

	msg1_cli := "Cliente " + strconv.Itoa(cl.id) + ": me corto el pelo"
	msg2_cli := "Cliente " + strconv.Itoa(cl.id) + ": termino de cortarme el pelo"
	msg3_cli := "Cliente " + strconv.Itoa(cl.id) + ": me siento en la sala de espera"
	msg4_cli := "Cliente " + strconv.Itoa(cl.id) + ": me voy de la barbería, está llena"

	recep.recepcionista <- cl
	msg_rec := <-cl.ch

	if msg_rec == "lleno" {
		fmt.Println(msg4_cli)
	} else {
		if msg_rec == "espera" {
			fmt.Println(msg3_cli)
			<-recep.sala_espera
			select {
			case b1.ch <- cl:
			case b2.ch <- cl:
			}
		}
		<-cl.ch
		fmt.Println(msg1_cli)
		<-cl.ch
		fmt.Println(msg2_cli)

	}
}

func monitor() {

	for {
		cl := <-recep.recepcionista

		if clientesesperando < MAXCLIENTESESPERANDO && barberosocup == MAXBARBEROS {
			clientesesperando++
			cl.ch <- "espera"
		} else if clientesesperando == MAXCLIENTESESPERANDO && barberosocup == MAXBARBEROS {
			cl.ch <- "lleno"
		} else {
			select {
			case b1.ch <- cl:
				cl.ch <- "OK"
			case b2.ch <- cl:
				cl.ch <- "OK"
			case recep.sala_espera <- "OK":
				clientesesperando--
			}
		}
	}
}

func main() {

	go monitor()

	go barbero(b1)
	go barbero(b2)

	for i := 0; i < 500; i++ {
		go cliente(i)
		time.Sleep(100 * time.Millisecond)
	}
}
