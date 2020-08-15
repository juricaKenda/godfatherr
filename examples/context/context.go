package main

import "github.com/juricaKenda/godfatherr"

const INTERNAL = "internal server"

func main() {
	err := someFunc()
	if err.ContainsCtx(INTERNAL) {
		err.Print()
	}
}
func someFunc() *godfatherr.Error {
	return godfatherr.New("usual err string").WithCtx(INTERNAL)
}
