BIN_DIR=_output/bin

REG = registry.cn-shanghai.aliyuncs.com
NS = charstal
TAG = v0.0.1
REPO = k8s-scheduler

.EXPORT_ALL_VARIABLES:

all: local

init:
	mkdir -p ${BIN_DIR}

local: init
	go build -o=${BIN_DIR}/scheduler-framework-demo ./cmd/scheduler

build-linux: init
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o=${BIN_DIR}/scheduler-framework-demo ./cmd/scheduler

image: build-linux
	docker build --no-cache . -t $(REG)/$(NS)/$(REPO):$(TAG)

push: image
	docker push $(REG)/$(NS)/$(REPO):$(TAG)

update:
	go mod download
	go mod tidy
	go mod vendor

clean:
	rm -rf _output/
	rm -f *.log
