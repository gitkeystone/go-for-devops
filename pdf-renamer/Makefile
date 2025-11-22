## 默认变量， 获取当前 go 的环境变量
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
CURRENT_DIR := $(shell basename $$PWD)

## build 在指定环境编译
build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o out/$(CURRENT_DIR)-$(GOOS)-$(GOARCH).exe .

## 通过指定环境变量， 执行交叉编译
buildx:
	GOOS=windows GOARCH=amd64 make build
## 清理编译结果
clean:
	rm -rf out/
