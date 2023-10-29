package main

import (
	"fmt"
	initModule "github.com/NeptuneYeh/golang_practice_01_NBA/init"
)

func main() {
	initStruct := initModule.NewMainInitProcess()
	initStruct.Run()
	fmt.Println("End main.go")
}
