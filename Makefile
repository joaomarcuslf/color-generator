GO_BUILD_ENV := CGO_ENABLED=0 GOOS=linux GOARCH=amd64

dev:
	go run main.go

test:
	go test -cover ./...

build:
	go build main.go

build-service:
	docker build . -t color-generator

start-service:
	docker run -d --name color-generator \
		-p 5003:5003 \
		--restart=always \
		color-generator

stop-service:
	docker stop color-generator
