build:
	docker-compose build
.PHONY: build

up:
	docker-compose up -d
.PHONY: up

down:
	docker-compose down
.PHONY: up

rebuild:
	docker-compose down
	docker-compose up -d --build
.PHONY: rebuild

down-full:
	docker-compose down
	docker rmi current-time-service metrics-service
.PHONY: down-full
