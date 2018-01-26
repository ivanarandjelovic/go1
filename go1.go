package main

import "runtime"
import "fmt"

func main() {
	var x = runtime.GOMAXPROCS(0)
	fmt.Println("Max procs = ", x)

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

func doSometime(kanal chan string, kraj chan bool) {
	for str := range kanal {
		println(str)
	}
	kraj <- true
}
