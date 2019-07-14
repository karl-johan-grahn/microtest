APP?=microtest
PORT?=8090
PROJECT?=github.com/karl-johan-grahn/microtest

RELEASE?=0.0.1
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')
CONTAINER_IMAGE?=docker.io/karljohangrahn/${APP}

GOOS?=linux
GOARCH?=amd64

clean:
	rm -f ${APP}

build: clean
	CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build \
		-ldflags "-s -w -X ${PROJECT}/version.Release=${RELEASE} \
		-X ${PROJECT}/version.Commit=${COMMIT} -X ${PROJECT}/version.BuildTime=${BUILD_TIME}" \
		-o ${APP}

container: build
	docker build -t $(CONTAINER_IMAGE):$(RELEASE) \
		--build-arg CREATED=$(BUILD_TIME) \
		--build-arg VERSION=$(RELEASE) \
		--build-arg REVISION=$(COMMIT) \
		.

run: container
	docker stop $(APP) || true && docker rm $(APP) || true
	docker run --name ${APP} -p ${PORT}:${PORT} --rm \
		-e "PORT=${PORT}" \
		$(CONTAINER_IMAGE):$(RELEASE)

test:
	go test -v -race -coverprofile=coverage.out ./...

push: container
	docker push $(CONTAINER_IMAGE):$(RELEASE)

coverage:
	go tool cover -html=coverage.out