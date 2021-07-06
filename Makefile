.PHONY: build run rundocker

GOIMAGE=golang:1.16-stretch
CASSANDRAIMAGE=cassandra:4.0

build:
	go build cmd/app.go

run:
	go run cmd/app.go

builddocker:
	docker run --rm -v ${PWD}:/app -w /app ${GOIMAGE} make build

rundocker: runcassandra
	docker run --rm -v ${PWD}:/app -w /app -p 8080:8180 ${GOIMAGE} make run

runcassandra:
	docker run --rm -d -p 9042:9042 ${CASSANDRAIMAGE}
