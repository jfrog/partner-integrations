.PHONY: usage
usage:
	@.scripts/_usage.sh

.PHONY: image
image:
	@echo "++ Building docker image..."
	docker build -t ${ARTIFACTORY_URL}/${ARTIFACTORY_DOCKER_REPO}/example-go:$(VERSION) . -f Dockerfile

.PHONY: push
deploy:
	@echo "++ Pushing docker image..."
	docker login ${ARTIFACTORY_URL}
	docker push ${ARTIFACTORY_URL}/${ARTIFACTORY_DOCKER_REPO}/example-go:$(VERSION)

.PHONY: build
build: export GOARCH=amd64
build: export CGO_ENABLED=0
build: export GO111MODULE=on
build: export GONOSUMDB="jfrog.com/*"
build:
	@echo "++ Building go binary..."
	jfrog rt go build -a --installsuffix cgo --ldflags="-s" -o bin/example-go --build-name=$(BUILD_NAME) --build-number=$(BUILD_NUMBER)

.PHONY: publish
publish: export GOARCH=amd64
publish: export CGO_ENABLED=0
publish: export GO111MODULE=on
publish: export GONOSUMDB="jfrog.com/*"
publish:
	@echo "++ Publishing..."
	jfrog rt bag ${BUILD_NAME} ${BUILD_NUMBER} --config=${HOME}/.jfrog/jira-cli.conf
	jfrog rt bce ${BUILD_NAME} ${BUILD_NUMBER}
	jfrog rt bp ${BUILD_NAME} ${BUILD_NUMBER}

.PHONY: tag
tag:
	@.scripts/tag.sh

.PHONY: clean
clean:
	rm -rf bin/example-go

.PHONY: swag
swag:
	swag init
