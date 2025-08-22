package grpc

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"text/template"
	"tools/command"
	. "tools/types"
	"tools/utils"
	"unicode"
)

func Generate(servicesPath string, services []Service) {
	servicesPath = strings.TrimRight(servicesPath, "/")
	servicesMap := makeServicesMap(services)
	servicesDeps := makeServicesDeps(services)
	for serviceName, deps := range servicesDeps {
		service := servicesMap[serviceName]
		for _, dep := range deps {
			serviceDep := servicesMap[dep]
			pathSource := fmt.Sprintf("%s/%s/contracts/", servicesPath, dep)
			pathClient := fmt.Sprintf("%s/%s/clients/%s", servicesPath, service.Name, dep)
			pathSourceMask := fmt.Sprintf("%s/%s/contracts//*.proto", servicesPath, dep)
			err := utils.MakeDir(pathClient)
			if err != nil {
				err = fmt.Errorf("Dir for GRPC don`t created and has error: %w", err)
				fmt.Println(err)
				continue
			}
			comGRPC := command.New("protoc").
				AddShortParam("I", pathSource).
				AddFullParam("go_out", pathClient).
				AddFullParam("go_opt=paths", "source_relative").
				AddFullParam("go-grpc_out", pathClient).
				AddFullParam("go-grpc_opt=paths", "source_relative").
				Argument(pathSourceMask)
			// add 'option go_package'
			for _, pathGRPC := range serviceDep.Contracts.GRPC.Files {
				fileGRPC := strings.ReplaceAll(pathGRPC, pathSource, "")
				name := fmt.Sprintf("go_opt=M%s", fileGRPC)
				value := fmt.Sprintf("%s/clients/%s", dep, dep)
				nameGRPC := fmt.Sprintf("go-grpc_opt=M%s", fileGRPC)
				comGRPC.AddFullParam(name, value)
				comGRPC.AddFullParam(nameGRPC, value)
			}
			result, err := exec.Command("bash", "-c", comGRPC.Build()).CombinedOutput()
			if err != nil {
				err = fmt.Errorf("Generation GRPC has error: %w, %s", err, result)
				fmt.Println(err)
				continue
			}
			fmt.Printf("For service '%s' was generated GRPC from service '%s'\n", service.Name, dep)
			err = generateTemplate("../templates/go_config.go.tmpl", serviceDep.Name, pathClient, "config")
			if err != nil {
				err = fmt.Errorf("Generation template has error: %w", err)
				fmt.Println(err)
				continue
			}
			fmt.Printf("For service '%s' was generated config to service '%s'\n", service.Name, dep)
			err = generateTemplate("../templates/go_client.go.tmpl", serviceDep.Name, pathClient, "client")
			if err != nil {
				err = fmt.Errorf("Generation template has error: %w", err)
				fmt.Println(err)
				continue
			}
			fmt.Printf("For service '%s' was generated client to service '%s'\n", service.Name, dep)
		}
	}
}

func generateTemplate(templatePath, service, target, targetName string) error {
	target = strings.TrimRight(target, "/")
	data := struct {
		Service      string
		ServiceCamel string
		ServiceUpper string
	}{
		Service:      service,
		ServiceCamel: UpperFirst(service),
		ServiceUpper: strings.ToUpper(service),
	}
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}
	targetFile := fmt.Sprintf("%s/%s_%s.go", target, service, targetName)
	outputFile, err := os.Create(targetFile)
	if err != nil {
		return err
	}
	defer outputFile.Close()
	err = tmpl.Execute(outputFile, data)
	return err
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

// UpperFirst capitalizes the first letter of a string.
func UpperFirst(s string) string {
	if s == "" {
		return ""
	}
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}
