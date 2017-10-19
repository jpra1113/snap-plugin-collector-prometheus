init:
	glide install

build-release:
	CGO_ENABLED=0 go build -a -installsuffix cgo

build-local:
	go build .
