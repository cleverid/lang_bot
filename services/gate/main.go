package main

import (
	"fmt"
	"gate/app"
	"gate/utils"
)

func main() {
	utils.LoadEnv(".env")
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
