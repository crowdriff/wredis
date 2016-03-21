version=0.1.5

.PHONY: all

all:
	@echo "make <cmd>"
	@echo ""
	@echo "commands:"
	@echo "  build         - build the dist binary"
	@echo "  clean         - clean the dist build"
	@echo "  deps          - pull and setup dependencies"
	@echo "  test          - run tests"
	@echo "  tools         - go get's a bunch of tools for development"
	@echo "  update_deps   - update deps glock file"

build: clean
	@go build ./...
	@go vet ./...
	@golint ./...

clean:
	@rm -rf ./bin

deps:
	@glock sync -n github.com/crowdriff/wredis < Glockfile

test:
	@ginkgo -r -v -cover -race

tools:
	go get github.com/robfig/glock
	go get github.com/golang/lint/golint
	go get github.com/onsi/ginkgo/ginkgo
	go get github.com/onsi/gomega

update_deps:
	@glock save -n github.com/crowdriff/wredis > Glockfile
