
.PHONY: start_dev_mysql
start_dev_mysql:
	docker run --name ggltest_mysql \
		-p 3306:3306 \
		--env-file ./.env \
		-v ${PWD}/data/mysql:/var/lib/mysql -d --rm mysql:5.7.39

.PHONY: gen
gen:
	go generate ./...

.PHONY: test
test: gen
	go test ./...

.PHONY: test_cover
test_cover: gen
	go test ./... -coverprofile cover.out
	go tool cover -html=cover.out