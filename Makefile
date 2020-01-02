pro:
	docker rmi -f web-service:1.0
	docker-compose up -d
local:
	go run main.go