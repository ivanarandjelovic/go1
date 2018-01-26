package main

import (
	"fmt"
	"runtime"

	"github.com/ivanarandjelovic/go1/struct1"
)

func main() {

	var x = runtime.GOMAXPROCS(0)
	fmt.Println("Max procs = ", x)

	//	goruts()

	structs()
}

func doSometime(kanal chan string, kraj chan bool) {
	for str := range kanal {
		println(str)
	}
	kraj <- true
}

func goruts() {
	var kanalica = make(chan string)
	var kraj = make(chan bool)

	go doSometime(kanalica, kraj)

	for i := 65; i < 95; i++ {
		println("mecem " + string(i))
		kanalica <- string(i)
	}
	close(kanalica)

	<-kraj
}

// structs plays with structures
func structs() {
	s := struct1.NewStr1(1, "mile")

	s.SayHello("Ivan")

	fmt.Println("Structure s is: ", s)
}
