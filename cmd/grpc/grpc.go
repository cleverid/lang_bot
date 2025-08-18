package grpc

import (
	"cmd/command"
	. "cmd/types"
	"cmd/utils"
	"fmt"
	"os/exec"
)

func Generate(servicesPath string, services []Service) {
	servicesMap := makeServicesMap(services)
	servicesDeps := makeServicesDeps(services)
	for serviceName, deps := range servicesDeps {
		service := servicesMap[serviceName]
		for _, dep := range deps {
			pathSource := fmt.Sprintf("%s/%s/contracts/", servicesPath, dep)
			pathClient := fmt.Sprintf("%s/%s/clients/%s", servicesPath, service.Name, dep)
			pathSourceMask := fmt.Sprintf("%s/%s/contracts//*.proto", servicesPath, dep)
			err := utils.MakeDir(pathClient)
			if err != nil {
				err = fmt.Errorf("Dir for GRPC don`t created and has error: %w", err)
				fmt.Println(err)
			}
			comGRPC := command.New("protoc")
			comGRPC.AddShortParam("I", pathSource)
			comGRPC.AddFullParam("go_out", pathClient)
			comGRPC.AddFullParam("go_opt=paths", "source_relative")
			comGRPC.AddFullParam("go-grpc_out", pathClient)
			comGRPC.AddFullParam("go-grpc_opt=paths", "source_relative")
			comGRPC.Argument(pathSourceMask)
			_, err = exec.Command("bash", "-c", comGRPC.Build()).CombinedOutput()
			if err != nil {
				err = fmt.Errorf("Generation GRPC has error: %w", err)
				fmt.Println(err)
			}
			fmt.Printf("For service '%s' was generated GRPC from service '%s'\n", service.Name, dep)
		}
	}
}

func makeServicesMap(services []Service) map[string]Service {
	servicesMap := make(map[string]Service)
	for _, service := range services {
		servicesMap[service.Name] = service
	}
	return servicesMap
}

func makeServicesDeps(services []Service) map[string][]string {
	servicesDeps := make(map[string][]string)
	for _, service := range services {
		servicesDeps[service.Name] = []string{}
		// Self
		if len(service.Contracts.GRPC.Files) > 0 {
			servicesDeps[service.Name] = append(servicesDeps[service.Name], service.Name)
		}
		// Deps
		for _, dep := range service.Clients {
			servicesDeps[service.Name] = append(servicesDeps[service.Name], dep.Service)
		}
	}
	return servicesDeps
}
