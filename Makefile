.PHONY: build run rundocker

GOIMAGE=golang:1.17-stretch
CASSANDRAIMAGE=cassandra:4.0
POSTGRESIMAGE=postgres:13.3

build:
	go build cmd/app.go

run:
	go run cmd/app.go

test:
	go test -v -count=1 -race -cover ./...

getgolangci:
	cd tools && go install github.com/golangci/golangci-lint/cmd/golangci-lint

lint: getgolangci
	golangci-lint run --config .golangci.yml

imports:
	goimports -l -w .

builddocker:
	docker run --rm -v ${PWD}:/app -w /app ${GOIMAGE} make build

lintdocker:
	docker run --rm -v ${PWD}:/app -w /app ${GOIMAGE} make lint

rundocker: runcassandra
	docker run --rm -v ${PWD}:/app -w /app -p 8080:8180 ${GOIMAGE} make run

runcassandra:
	docker run --rm -d -p 9042:9042 ${CASSANDRAIMAGE}

runpostgres:
	docker run --rm -d -p 5432:5432 -e POSTGRES_PASSWORD=pass -v ${PWD}/docker/postgres/initdb:/docker-entrypoint-initdb.d ${POSTGRESIMAGE}

buildeb:
	GOARCH=amd64 GOOS=linux go build -o build/bin/application cmd/app.go
	# on docker case
	apt update && apt install -y zip
	cd build && zip -r ../app.zip *

buildforerasticbeans:
	docker run --rm -v ${PWD}:/app -w /app ${GOIMAGE} make buildeb
