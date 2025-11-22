## 默认变量， 获取当前 go 的环境变量
CGO_ENABLED ?= 0
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
LDFLAGS := -s -w
CURRENT_DIR := $(shell basename $$PWD)

## build 在指定环境编译
build:
	CGO_ENABLED=$(CGO_ENABLED) GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags="$(LDFLAGS)" -o out/$(CURRENT_DIR)-$(GOOS)-$(GOARCH)$(shell [ "$(GOOS)" = "windows" ] && echo ".exe") .

## 通过指定环境变量， 执行交叉编译
buildx:
	GOOS=linux GOARCH=amd64 make build
	GOOS=linux GOARCH=arm64 make build
	GOOS=darwin GOARCH=amd64 make build
	GOOS=darwin GOARCH=arm64 make build
	GOOS=windows GOARCH=amd64 make build
	GOOS=windows GOARCH=arm64 make build

## 仅编译 Windows 版本
build-windows:
	GOOS=windows GOARCH=amd64 make build
	GOOS=windows GOARCH=arm64 make build

## 清理编译结果
clean:
	rm -rf out/

.PHONY: build buildx build-windows clean
