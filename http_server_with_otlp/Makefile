install:
	@go install github.com/swaggo/swag/cmd/swag@latest

swagger-gen:
	@swag init

containers:
	docker-compose -f ./.containers/docker-compose.yml up -d

download: containers
	go mod download

run: download
	GO_ENV=local go run main.go api