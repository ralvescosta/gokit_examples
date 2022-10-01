containers:
	docker-compose -f ./.containers/docker-compose.yml up -d

download: containers
	go mod download

api: download
	GO_ENV=local go run main.go api

consumer: download
	GO_ENV=local go run main.go consumer
