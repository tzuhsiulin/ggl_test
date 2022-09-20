
.PHONY: start_dev_mysql
start_dev_mysql:
	docker run --name ggltest_mysql \
		-p 3306:3306 \
		--env-file ./.env \
		-v ${PWD}/data/mysql:/var/lib/mysql -d --rm mysql:5.7.39


gen:
	go generate ./...