package grpc

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

var SERVICES_PATH = "../services"

func Generate() {
	services, err := loadServices(SERVICES_PATH)
	fmt.Println(services)
	if err != nil {
		err = fmt.Errorf("Load services from %s has error: %w", SERVICES_PATH, err)
		log.Fatalln(err)
	}
}

func loadServices(dir string) (services []Service, err error) {
	services = []Service{}
	items, err := os.ReadDir(dir)
	if err != nil {
		err = fmt.Errorf("Read %s has error: %w", SERVICES_PATH, err)
		return nil, err
	}
	for _, item := range items {
		if !item.IsDir() {
			continue
		}
		service, err := loadService(item)
		if err != nil {
			return nil, err
		}
		services = append(services, *service)
	}
	return
}

func loadService(dir os.DirEntry) (service *Service, err error) {
	pathServiceJSON := fmt.Sprintf("%s/%s/service.json", SERVICES_PATH, dir.Name())
	pathContracts := fmt.Sprintf("%s/%s/contracts/", SERVICES_PATH, dir.Name())
	service, err = loadServiceJSON(pathServiceJSON)
	if err != nil {
		return nil, err
	}
	contracts, err := loadFilesByPathAndExtension(pathContracts, "proto")
	if err != nil {
		return nil, err
	}
	service.Contracts.GRPC.Files = contracts
	return service, nil
}

func loadServiceJSON(path string) (service *Service, err error) {
	file, err := os.ReadFile(path)
	if err != nil {
		err = fmt.Errorf("Read %s file has error: %w", path, err)
		return nil, err
	}
	service = &Service{}
	err = json.Unmarshal(file, service)
	if err != nil {
		err = fmt.Errorf("Parsing %s file has error: %w", path, err)
		return nil, err
	}
	return
}

func loadFilesByPathAndExtension(dir string, ext string) (files []string, err error) {
	dir = strings.TrimRight(dir, "/")
	files = []string{}
	items, err := os.ReadDir(dir)
	if err != nil {
		return files, nil
	}
	for _, item := range items {
		if !strings.Contains(item.Name(), ext) {
			continue
		}
		path := fmt.Sprintf("%s/%s", dir, item.Name())
		files = append(files, path)
	}
	return
}
