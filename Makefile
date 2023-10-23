up:
	docker compose up -d
build:
	docker build . -t gemanueldev/utils-api-go
buildUp:
	docker build . -t gemanueldev/utils-api-go
	docker-compose up -d
buildUpLog:
	docker build . -t gemanueldev/utils-api-go
	docker-compose up