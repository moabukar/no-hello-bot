
go/install:
	go get -v

go/tidy:
	go mod tidy -compat=1.19


go/test:
	go test -v --cover ./...

go/lint:
	golangci-lint run --allow-parallel-runners

go/run:
	./scripts/go.sh run

go/build:
	./scripts/go.sh build

# dev/up:
# 	./scripts/dev.sh up

# dev/destroy:
# 	./scripts/dev.sh destroy