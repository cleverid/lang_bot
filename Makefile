complete -W "\`grep -oE '^[a-zA-Z0-9_.-]+:([^=]|$)' ?akefile | sed 's/[^a-zA-Z0-9_.-]*$//'\`" make

help:
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

init: ## Init
	docker network create -d bridge lang-bot 

up: infra-up ## Up
down: infra-down ## Down
stop: infra-stop ## Stop
ps: ## View all containers
	docker ps | grep lang-bot-
fmt: fmt-gate fmt-user ## Format
test: ## Test
	go test ./... 
gen: ## Geneate 
	./grpc.sh

fmt-gate: ## Format gate
	cd services/gate && go fmt ./...
fmt-user: ## Format user
	cd services/user && go fmt ./...
 
dev-gate: ## Develop mod for gate service
	cd services/gate && go run main.go 
dev-user: ## Develop mod for user service
	cd services/user && go run main.go 

infra-up: redis-up postgres-up ## Infra. Up
infra-down: redis-down postgres-down ## Infra. Down
infra-stop: redis-stop postgres-stop ## Infra. Stop

