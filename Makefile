APP_NAME :=oscar-service
VERSION := $(shell git describe --tags --abbrev=0)
VERSION:=1.0
DOCKER_USER=eneoti
DOCKER_REPO=754404031763.dkr.ecr.ap-southeast-1.amazonaws.com
HELM := $(shell helm ls|grep '$(APP_NAME)')
ifeq ($(VERSION),)
	VERSION:= "v1.0"
endif
# HELP
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help


# DOCKER TASKS
# Build the container
build-docker: ## Build the container
	docker build --build-arg APP_NAME=$(APP_NAME) -t $(APP_NAME) .
clear-none:
	docker rmi -f `docker images -a |grep 'none'|awk '{print \$$3}'`

clear:
ifeq ($(HELM),)
	echo 'not exist broker'
else
	--helm delete $(APP_NAME) --purge
endif
	--docker rmi -f `docker images -a |grep '$(APP_NAME)'|awk '{print\$$3}'`
	# --docker rmi -f `docker images -f "dangling=true" -q `

clearall:
	docker rmi -f `docker images -a |grep '$(APP_NAME)\|none'|awk '{print\$$3}'`

clean: stop clear

proto:
	./pkg/grpc/proto/compile.sh
build-nc: ## Build the container without caching
	docker build --no-cache -t $(APP_NAME) .

run: ## Run container
	docker run \
	--name $(APP_NAME) -p 8080:8080  -d $(APP_NAME)

up: build run ## Run container on
stop: ## Stop and remove a running container
	docker stop $(APP_NAME); docker rm -f $(APP_NAME)

stopall:
	docker rm -f `docker ps -a -q`

run-go:
	go run cmd/api/main.go

prebuild:
	GOOS=linux go build -o plugins/build/alerts/equal plugins/alerts/equal.go

build-go:
	GOOS=linux go build -o $(APP_NAME) cmd/api/main.go
	# go build -o iot-service cmd/api/main.go
swagger:
	swagger generate spec -m -b ./pkg/api -o ./assets/swaggerui/swagger.json
swagger-internal:
	swagger generate spec -m -b ./pkg/intercom -o ./assets/swaggerui/swagger-internal.json

release: build-nc publish clear ## Make a release by building and publishing the `{version}` ans `latest` tagged containers to ECR

# Docker publish
publish: login publish-latest ## Publish the `{version}` ans `latest` tagged containers to ECR

publish-latest: tag-latest ## Publish the `latest` taged container to ECR
	--aws ecr create-repository --repository-name $(APP_NAME)
	--aws ecr batch-delete-image --repository-name $(APP_NAME) --image-ids imageTag=latest
	--aws ecr batch-delete-image --repository-name $(APP_NAME) --image-ids imageTag=1.0
	@echo 'publish latest to $(DOCKER_REPO)'
	docker push $(DOCKER_REPO)/$(APP_NAME):latest

publish-version: tag-version ## Publish the `{version}` taged container to ECR
	@echo 'publish $(VERSION) to $(DOCKER_REPO)'
	docker push $(DOCKER_REPO)/$(APP_NAME):$(VERSION)

# Docker tagging
tag: tag-latest tag-version ## Generate container tags for the `{version}` ans `latest` tags

tag-latest: ## Generate container `{version}` tag
	@echo 'create tag latest'
	docker tag $(APP_NAME) $(DOCKER_REPO)/$(APP_NAME):latest

tag-version: ## Generate container `latest` tag
	@echo 'create tag $(VERSION)'
	docker tag $(APP_NAME) $(DOCKER_REPO)/$(APP_NAME):$(VERSION)

# login to AWS-ECR
login: ## login dockerhub
	$(shell aws ecr get-login --no-include-email --region ap-southeast-1)

deploy:
	helm install --name $(APP_NAME) deployment --set ENV=prod \
	--set entryPoint=$(APP_NAME) \
	--set serverName=$(APP_NAME) \
	--set image.repository=$(DOCKER_REPO)/$(APP_NAME) \
	--set fullnameOverride=$(APP_NAME)


update-deploy:
	--helm upgrade --name $(APP_NAME) deployment --set ENV=prod

undeploy:
	--helm delete $(APP_NAME) --purge
port-forward:
	kubectl -n iot port-forward  svc/$(APP_NAME)-production 8080:8080
version: ## Output the current version
	@echo $(VERSION)
commit:
	git add  .
	--git commit -m "$m"
commitnPush: commit
	--git push origin master
ungittag:
	echo $(VERSION)
	--git push --delete origin $(VERSION)
	--git tag --delete $(VERSION)
gittag:
	git tag -a "$(VERSION)" -m "$(VERSION)"
	git push --tags
quicktag: ungittag commit gittag
