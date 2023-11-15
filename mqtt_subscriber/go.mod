module github.com/ralvescosta/gokit_example/mqtt_subscriber

go 1.21.4

require (
	github.com/ralvescosta/gokit/configs_builder v1.15.0
	github.com/ralvescosta/gokit/logging v1.15.0
	github.com/ralvescosta/gokit/mqtt v0.0.0-20231115222828-1857f3bf6e25
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/eclipse/paho.mqtt.golang v1.4.3 // indirect
	github.com/gorilla/websocket v1.5.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/ralvescosta/dotenv v1.0.4 // indirect
	github.com/ralvescosta/gokit/configs v1.15.0 // indirect
	github.com/stretchr/objx v0.5.1 // indirect
	github.com/stretchr/testify v1.8.4 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.26.0 // indirect
	golang.org/x/net v0.18.0 // indirect
	golang.org/x/sync v0.5.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/ralvescosta/gokit/configs => ../../gokit/configs

replace github.com/ralvescosta/gokit/configs_builder => ../../gokit/configs_builder
