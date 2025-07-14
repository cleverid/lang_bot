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
dev-lang: ## Develop mod for lang service
	go run cmd/lang/lang.go 
dev-user: ## Develop mod for lang service
	go run cmd/user/user.go 
fmt: ## Format
	go fmt ./... 
test: ## Test
	go test ./... 

infra-up: redis-up postgres-up ## Infra. Up
infra-down: redis-down postgres-down ## Infra. Down
infra-stop: redis-stop postgres-stop ## Infra. Stop

