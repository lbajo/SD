package logiclog

import (
    "os"
    "testing"
)

func newFile() *os.File{
    f, err := os.Create("/tmp/LogiclogTest")
    if err != nil{
        panic(err)
    }
    return f
}

func TestLogiclog(t *testing.T){
    file := newFile()
    log1 := NewLog("Lorena",file)
    log2 := NewLog("Nuria",file)
    WriteLog(log1,"Hola soy Lorena")
    WriteLog(log2,"Hola soy Nuria")
    ReceiveLog(log2, SendLog(log1, "Qué haces Nuria?"))
    ReceiveLog(log1, SendLog(log2, "Nada, vamos a comer"))
    ReceiveLog(log2, SendLog(log1, "Vale"))
    WriteLog(log2,"Voy a apagar el ordenador")
}
