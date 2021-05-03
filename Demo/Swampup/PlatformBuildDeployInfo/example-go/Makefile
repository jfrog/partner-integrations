.PHONY: usage
usage:
	@.scripts/_usage.sh

.PHONY: image
image:
	@echo "++ Building docker image..."
	docker build -t swampup.jfrog.tech/docker-example-virtual/example-go:$(VERSION) . -f Dockerfile

.PHONY: deploy
deploy:
	@echo "++ Pushing docker image..."
	docker login swampup.jfrog.tech
	docker push swampup.jfrog.tech/docker-example-virtual/example-go:$(VERSION)

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
	@echo "++ Building go binary..."
	jfrog rt go-publish $(GO_VIRTUAL_REPO) v$(VERSION) --build-name=$(BUILD_NAME) --build-number=$(BUILD_NUMBER)

.PHONY: tag
tag:
	@.scripts/tag.sh

.PHONY: clean
clean:
	rm -rf bin/example-go

.PHONY: swag
swag:
	swag init
