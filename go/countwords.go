//LORENA BAJO REBOLLO		TELEMÁTICA
package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
)

var m = make(map[string]int)

func imprimir() {
	var keys []string

	sort.Strings(keys)
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		fmt.Println(key+":", m[key])
	}
}

func procesarpalabra(palabra string) {

	valor, ok := m[palabra]

	if !ok {
		m[palabra] = 1
	} else {
		m[palabra] = valor + 1
	}

}

func procesarfichero(fich string) {
	data, err := os.Open(fich)

	if err != nil {
		fmt.Println(errors.New("Error al leer el fichero"))
	}

	scanner := bufio.NewScanner(data)

	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		palabra := scanner.Text()
		procesarpalabra(palabra)
	}
}

func main() {

	args := os.Args[1:]
	num_args := len(args)

	for i := 1; i <= num_args; i++ {
		procesarfichero(os.Args[i])
	}

	imprimir()
}
