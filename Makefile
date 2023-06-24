-include .env

.PHONY: build

# build app
build: build-app

build-app:
	@echo " > Building [aquafarm app]..."
	@cd ./cmd/http/ && go build -o ../../bin/http && cd ..
	@echo " > Finished building [aquafarm app]"

# run app
run-app: build-app
	@echo " > Building [aquafarm app]..."
	@ ./bin/http
	@echo " > Finished building [aquafarm app]"
