.DEFAULT_GOAL := build

fmt:
	go fmt ./...
.PHONY:fmt

lint: fmt
	golint ./...
.PHONY:lint

vet: fmt
	go vet ./...
.PHONY:vet

build: vet
	go build .
.PHONY:build

build_arm_7:
	env GOOS=linux GOARCH=arm GOARM=7 go build -o fuel-prices_arm-v7 fuel-prices.go
.PHONY:build_arm_7
