package main

import (
	"cmd/grpc"
	"cmd/services"
	"fmt"
	"log"
)

const SERVICES_PATH = "../services/"

func main() {
	services, err := services.LoadServices(SERVICES_PATH)
	if err != nil {
		err = fmt.Errorf("Load services has error: %w", err)
		log.Fatal(err)
	}
	grpc.Generate(SERVICES_PATH, services)
	if err != nil {
		err = fmt.Errorf("Generate GRPC has error: %w", err)
		log.Fatal(err)
	}
}
