dep:
	go mod tidy

docker_build:
	docker build -t micheltank/price-alert-service ./price-alert-service
	docker build -f price-alert-job/Dockerfile -t micheltank/price-alert-job .
	docker build -t micheltank/notification-service ./notification-service

docker_compose:
	docker-compose up -d

all: docker_build docker_compose