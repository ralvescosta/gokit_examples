containers:
	docker-compose -f ./.containers/docker-compose.yml up

download: containers
	go mod download

run: download
	GO_ENV=local go run main.go api