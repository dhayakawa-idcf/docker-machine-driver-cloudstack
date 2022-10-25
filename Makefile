all: build

build:
	CGO_ENABLED=0 go build -ldflags="-w -s" -o docker-machine-driver-cloudstack cmd/docker-machine-driver-cloudstack/main.go
	chmod +x docker-machine-driver-cloudstack
