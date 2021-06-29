# .PHONY: help
help: # Show this help
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

MSG=$(shell git show -s --format=%s)
TAG=$(shell /bin/date "+%Y%m%d%H%M%S")

init:
	go mod init exponea.com
	go mod tidy

run: ## run
	go run main.go

docker: ## docker
	docker build -t exponea .
	docker run -p 8080:8080 exponea:latest
