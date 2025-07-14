package main

import (
	"fmt"
	"lb/services/lang/app"
	"lb/utils"
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
