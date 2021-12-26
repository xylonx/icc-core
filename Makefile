PROJECT:=icc-core
CONFIG_FILE:=config.local.yaml
GO_MODULE_STATE:=$(shell go env GO111MODULE)
GO_PROXY:=$(shell go env GOPROXY)

.PHONY: build dep serve clean

.DEFAULT: build

build: dep
	go build -o ${PROJECT}

${PROJECT}: build

dep:
ifneq ($(GO_MODULE_STATE), on)
	go env -w GO111MODULE="on"
endif
ifeq ($(GO_PROXY), https://proxy.golang.org,direct)
	go env -w GOPROXY="https://goproxy.cn,direct"
endif
	go mod tidy

server: build
	./${PROJECT} -c ${CONFIG_FILE}

clean:
	rm -rf ${PROJECT}