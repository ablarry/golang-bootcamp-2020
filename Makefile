help:
	@echo "Commands:"
	@echo "\tmake install\t Install dependencies."
	@echo "\tmake test\t Runit tests"
	@echo "\tmake run\t Run"
	@echo "\tmake linter\t Linter of code project"
	@echo "\tmake coverage\t Coverage of code project"
	@echo "\tmake mocks\t Create code mocks"

install:
	@echo "Make: install"
	go mod download


linter:
	@echo "Make: linter"
	go get -u golang.org/x/lint/golint
	golint ./pkg/...
	golint ./cmd/pokemon/main.go

coverage:
	@echo "Make: coverage"
	go clean -testcache && go test  -coverprofile=coverage.out ./...
	go tool cover  -html=coverage.out

test:
	@echo "Make: test"
	go clean -testcache && go test  -coverprofile=coverage.out ./...

run:
	@echo "Make: run"
	go run ./cmd/pokemon/main.go

mocks:
	@echo "Make: mocks"
	mockgen -source=./pkg/repo/repo.go -destination=./pkg/repo/mockrepo.go -package=repo
	mockgen -source=./pkg/service/service.go -destination=./pkg/service/mockservice.go -package=service


