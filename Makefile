GOCMD=go
BUILD_TAGS=-tags netgo
LDFLAG=-ldflags="-s -w"
GOBUILD=env GOOS=linux GOARCH=amd64 $(GOCMD) build $(BUILD_TAGS) $(LDFLAG)
BIN_DIRECTORY=bin

.PHONY: start_dev_mysql
start_dev_mysql:
	docker run --name ggltest_mysql \
		-p 3306:3306 \
		--env-file ./.env \
		-v ${PWD}/data/mysql:/var/lib/mysql -d --rm mysql:5.7.39

.PHONY: test
test:
	go generate ./...
	go test ./...

.PHONY: test_cover
test_cover:
	go generate ./...
	go test ./... -coverprofile cover.out
	go tool cover -html=cover.out

bin/api:
	 $(GOBUILD) -o $(BIN_DIRECTORY)/api -v cmd/api/main.go

bin/db_migration:
	$(GOBUILD) -o $(BIN_DIRECTORY)/db_migration -v cmd/db_migration/main.go

bin: bin/api bin/db_migration

.PHONY: clean
clean:
	rm -rf bin