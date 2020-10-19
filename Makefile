BIN_DIR=_output/bin

# If tag not explicitly set in users default to the git sha.
# TAG ?= ${shell (git describe --tags --abbrev=14 | sed "s/-g\([0-9a-f]\{14\}\)$/+\1/") 2>/dev/null || git rev-parse --verify --short HEAD}

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
	docker build --no-cache . -t $(NS)/$(REPO):$(TAG)

push: image
	docker push $(NS)/$(REPO):$(TAG)

update:
	go mod download
	go mod tidy
	go mod vendor

clean:
	rm -rf _output/
	rm -f *.log