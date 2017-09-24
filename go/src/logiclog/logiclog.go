package logiclog

import (
	"bufio"
	"errors"
	"fmt"
	"logicclock"
	"math"
	"os"
	"sort"
)

type Log struct {
	Cl   logicclock.Clock
	file *os.File
}

type Msg struct {
	txt string
	cl  string
}

type list struct {
	ck []logicclock.Clock
}

var listClocks = new(list)

func NewLog(name string, file *os.File) Log {

	log := Log{logicclock.NewClock(name), file}
	return log
}

func WriteLog(log Log, txt string) {

	logicclock.AddClock(log.Cl)
	log.Cl.Log = txt
	js := logicclock.ToJSON(log.Cl)
	log.file.WriteString(js + "\n")
}

func SendLog(log Log, txt string) Msg {

	WriteLog(log, "Enviado: "+txt)
	js := logicclock.ToJSON(log.Cl)
	msg := Msg{txt, js}
	return msg
}

func ReceiveLog(log Log, msg Msg) {
	clock := logicclock.FromJSON(msg.cl)
	logicclock.Max(log.Cl, clock)
	WriteLog(log, "Recibido: "+msg.txt)
}

func PlusLog(log Log) {
	logicclock.AddClock(log.Cl)
}

func GetMsg(msg Msg) string {
	return msg.txt
}

/*
func ProcessLines(line string, file string) Line{
	newLine := Line{}
	text := ""
	text2 := ""
	text3 := ""
	text4 := ""
	text5 := ""
	text6 := ""

	newLine.mark = logicclock.NewClock(file)
	logicclock.FromJSON(newLine.mark,line)

	newLine.text = text

	words2 := strings.Split(line, " ")
	text2 =  words2[len(words2)-1]
	fmt.Println("22222 ", text2)

	words3 := strings.Split(text2, "Map")
	text3 =  words3[len(words3)-1]
	fmt.Println("33333 ", text3)

	words4 := strings.Split(text3, ":{")
	text4 =  words4[len(words4)-1]
	fmt.Println("44444 ", text4)


	words5 := strings.Split(text4, "}}")
	text5 =  words5[0]
	fmt.Println("55555 ", text5)

	text6 = "{"+text5+"}"

	fmt.Println("6666 ", text6)

	return newLine

}*/

func ProcessFiles(file string) {

	data, err := os.Open(file)

	if err != nil {
		fmt.Println(errors.New("Error al leer el fichero"))
	}

	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		cl := logicclock.FromJSON(scanner.Text())
		listClocks.ck = append(listClocks.ck, cl)
	}
}

func PrintLogs() {
	for _, c := range listClocks.ck {
		fmt.Println(logicclock.ToJSON(c))
	}
}

func (list list) Len() int {
	return len(list.ck)
}

func GetValue(list list, n int) float64 {
	var total float64 = 0
	for _, val := range list.ck[n].Map {
		total += float64(val) * float64(val)
	}

	return math.Sqrt(total)
}

func (list list) Less(i, j int) bool {

	v1 := GetValue(list, i)
	v2 := GetValue(list, j)

	if v1 < v2 {
		return true
	}

	return false
}

func (list list) Swap(i, j int) {
	list.ck[i], list.ck[j] = list.ck[j], list.ck[i]
}

func Order() {
	sort.Sort(listClocks)
	PrintLogs()
}
