PROJECT_NAME=valexer
GO_BUILD_ENV=CGO_ENABLED=0 GOOS=linux GOARCH=amd64
DOCKER_DIR=deployment/docker

.SILENT:
all: build docker compose

build:
	$(GO_BUILD_ENV) go build -o $(PROJECT_NAME) main.go

compose:
	cd $(DOCKER_DIR); \
	docker-compose up;

clean:
	docker system prune --volumes -f

docker: build
	mv $(PROJECT_NAME) $(DOCKER_DIR)/$(PROJECT_NAME); \
	cp -r web/ $(DOCKER_DIR)/web/; \
	cd $(DOCKER_DIR); \
	docker build -t $(PROJECT_NAME):latest .; \
	rm -rf $(PROJECT_NAME); \
	rm -rf web;