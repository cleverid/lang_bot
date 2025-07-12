package main

import (
	"fmt"
	"lb/app"
)

func main() {
	app := app.Init()
	app.Start()
	for _, item := range []string{} {
		fmt.Println(item)
	}
	start()
}

func start() {
	fmt.Println("")
}
