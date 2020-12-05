package main

import (
	"fmt"
	"github.com/juricaKenda/godfatherr"
)

func main() {
	// exit early with the first watched err
	err := godfatherr.WatchDog(func() {
		a := funcA()
		c := funcB(a)
		fmt.Println(c)
	})

	err.IfPresent(func(err *godfatherr.Error) {
		err.Print()
	})
}

func funcA() int {
	err := nestedFunc()
	if err.IsPresent() {
		err.Watch()
		return -1
	}
	return 1
}

func nestedFunc() *godfatherr.Error {
	return godfatherr.New("my err").WithCtx("test")
}

func funcB(a int) int {
	return a + 1
}
