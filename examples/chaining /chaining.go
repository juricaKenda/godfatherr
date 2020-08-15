package main

import (
	"fmt"
	"github.com/juricaKenda/godfatherr"
)

func main() {
	funcOne().IfAbsent(funcTwo).IfAbsent(funcThree).IfPresent(log)
}

func funcOne() *godfatherr.Error {
	fmt.Println("executing func one")
	return godfatherr.Empty()
}

func funcTwo() *godfatherr.Error {
	fmt.Println("executing second func")
	return godfatherr.New("something went wrong")
}

func funcThree() *godfatherr.Error {
	return godfatherr.Empty()
}

func log(err *godfatherr.Error) {
	err.Print()
}
