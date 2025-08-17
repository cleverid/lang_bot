package grpc

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	dir := "../services"

	items, err := os.ReadDir(dir)
	if err != nil {
		log.Fatalln(err)
	}

	for _, item := range items {
		if !item.IsDir() {
			continue
		}
		servicePath := fmt.Sprintf("%s/%s/service.json", dir, item.Name())
		serviceFile, err := os.ReadFile(servicePath)
		if err != nil {
			log.Fatalln(err)
		}
		service := Service{}
		err = json.Unmarshal(serviceFile, &service)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(service)
	}
}
