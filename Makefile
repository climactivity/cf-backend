.PHONY: serve

.DEFAULT_GOAL := all

MAIN := ./cms
DOCKER_TAG := ghcr.io/climactivity/climate-friday-backend:latest


all: dependencies frontend serve

check-compiles: dependencies frontend build-go

build-go: 
	go build $(MAIN) 

dependencies: 
	npm install

frontend:
	npm run build 

serve: 
	go run $(MAIN) serve

collections: 
	go run $(MAIN) migrate collections

run-container:
	docker-compose up --build -d

build-container: 
	docker build . -f ./Dockerfile --tag $(DOCKER_TAG)

push-container: build-container
	docker push $(DOCKER_TAG)
