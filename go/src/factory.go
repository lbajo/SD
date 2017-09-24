package main

import (
	"fmt"
	"sem"
	"time"
)

type prodCons struct {
	id     int
	buff   map[int]int
	prod   *sem.Sem
	consum *sem.Sem
	pos    int
}

var (
	cables    = prodCons{0, make(map[int]int), sem.NewSem(0), sem.NewSem(1), 0}
	pantallas = prodCons{0, make(map[int]int), sem.NewSem(0), sem.NewSem(1), 0}
	carcasas  = prodCons{0, make(map[int]int), sem.NewSem(0), sem.NewSem(1), 0}
	placas    = prodCons{0, make(map[int]int), sem.NewSem(0), sem.NewSem(1), 0}
)

func producirPantallas() {
	for i := 0; i < 500; i++ {
		pantallas.consum.Down()
		pantallas.id = i
		pantallas.buff[i] = i
		pantallas.prod.Up()
		time.Sleep(50 * time.Millisecond)
	}
}

func producirCarcasas() {
	for i := 0; i < 500; i++ {
		carcasas.consum.Down()
		carcasas.id = i
		carcasas.buff[i] = i
		carcasas.prod.Up()
		time.Sleep(60 * time.Millisecond)
	}
}

func producirPlacas() {
	for i := 0; i < 500; i++ {
		placas.consum.Down()
		placas.id = i
		placas.buff[i] = i
		placas.prod.Up()
		time.Sleep(70 * time.Millisecond)
	}
}

func producirCables() {
	for i := 0; i < 500; i++ {
		cables.consum.Down()
		cables.id = i
		cables.buff[i] = i
		cables.prod.Up()
		time.Sleep(80 * time.Millisecond)
	}
}

func obtenerPantallas(idobtenido []int) int {
	count := 0
	pantallas.prod.Down()
	//idobtenido[5] = pantallas.id
	idobtenido[5] = pantallas.buff[pantallas.pos]
	pantallas.consum.Up()
	pantallas.pos++
	count++
	return count
}

func obtenerCarcasas(idobtenido []int) int {
	count := 0
	carcasas.prod.Down()
	//idobtenido[6] = carcasas.id
	idobtenido[6] = carcasas.buff[carcasas.pos]
	carcasas.consum.Up()
	carcasas.pos++
	count++
	return count
}

func obtenerPlacas(idobtenido []int) int {
	count := 0

	placas.prod.Down()
	//idobtenido[7] = placas.id
	idobtenido[7] = placas.buff[placas.pos]
	placas.consum.Up()
	placas.pos++
	count++
	return count
}

func obtenerCables(idobtenido []int) int {
	count := 0

	for i := 0; i < 5; i++ {
		cables.prod.Down()
		//idobtenido[i] = cables.id
		idobtenido[i] = cables.buff[cables.pos]
		cables.consum.Up()
		cables.pos++
		count++
	}

	return count
}

func obtener(idobtenido []int) int {
	cabl := obtenerCables(idobtenido)
	pant := obtenerPantallas(idobtenido)
	carc := obtenerCarcasas(idobtenido)
	placas := obtenerPlacas(idobtenido)

	if (cabl == 5) && (pant == 1) && (carc == 1) && (placas == 1) {
		return 0
	}

	return -1
}

func robot(idrobot int) {

	idobtenido := make([]int, 8)

	for {

		ok := obtener(idobtenido)

		if ok == 0 {

			fmt.Println("robot", idrobot, ", cables", idobtenido[0], idobtenido[1], idobtenido[2], idobtenido[3], idobtenido[4], "pantalla", idobtenido[5], "carcasa", idobtenido[6], "placa", idobtenido[7], "Comenzando")

			time.Sleep(200 * time.Millisecond)

			fmt.Println("robot", idrobot, ", cables", idobtenido[0], idobtenido[1], idobtenido[2], idobtenido[3], idobtenido[4], "pantalla", idobtenido[5], "carcasa", idobtenido[6], "placa", idobtenido[7], "Terminado")
		}
	}

}

func main() {

	go producirPantallas()
	go producirCarcasas()
	go producirCables()
	go producirPlacas()

	for i := 1; i < 4; i++ {
		go robot(i)
	}

	time.Sleep(58000 * time.Millisecond)
}
