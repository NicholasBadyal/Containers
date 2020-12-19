all: install fmt

deploy: image dock

install:
	go install -v

fmt:
	go fmt ./...

image:
	docker build -t nbadyal/simple-go-server:latest .

dock:
	docker run -it -p 5050:5050 nbadyal/simple-go-server

.PHONY: install fmt local