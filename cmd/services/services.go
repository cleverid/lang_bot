package services

import (
	. "cmd/types"
	"cmd/utils"
	"encoding/json"
	"fmt"
	"os"
)

func LoadServices(servicesPath string) (services []Service, err error) {
	services = []Service{}
	items, err := os.ReadDir(servicesPath)
	if err != nil {
		err = fmt.Errorf("Read %s has error: %w", servicesPath, err)
		return nil, err
	}
	for _, item := range items {
		if !item.IsDir() {
			continue
		}
		service, err := loadService(servicesPath, item)
		if err != nil {
			return nil, err
		}
		services = append(services, *service)
	}
	return
}

func loadService(servicesPath string, dir os.DirEntry) (service *Service, err error) {
	pathServiceJSON := fmt.Sprintf("%s/%s/service.json", servicesPath, dir.Name())
	pathContracts := fmt.Sprintf("%s/%s/contracts/", servicesPath, dir.Name())
	service, err = loadServiceJSON(pathServiceJSON)
	if err != nil {
		return nil, err
	}
	contracts, err := utils.LoadFilesByPathAndExtension(pathContracts, "proto")
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
