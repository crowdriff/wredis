version=0.0.2

.PHONY: all

# To cross compile for linux on mac, build go-linux cross compiler first using
# cd /usr/local/go/src
# sudo GOOS=linux GOARCH=amd64 CGO_ENABLED=0 ./make.bash --no-clean

all:
	@echo "make <cmd>"
	@echo ""
	@echo "commands:"
	@echo "  build         - build the dist binary"
	@echo "  clean         - clean the dist build"
	@echo "  install       - run go install for all sub packages"
	@echo "  test          - standard go test"
	@echo ""
	@echo "  tools         - go get's a bunch of tools for dev"
	@echo "  deps          - pull and setup dependencies"
	@echo "  update_deps   - update deps lock file"

build: clean
	@go build ./...
	@go vet ./...
	@golint ./...

clean:
	@rm -rf ./bin

coverage:
	@go test -cover -v ./...

deps:
	@glock sync -n github.com/crowdriff/wredis < Glockfile

install:
	@go install ./...

test:
	@ginkgo

tools:
	go get github.com/robfig/glock
	go get github.com/golang/lint/golint

update_deps:
	@glock save -n github.com/crowdriff/wredis > Glockfile
