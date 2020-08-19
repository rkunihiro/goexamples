.PHONY: all tools clean fmt fix lint test build

TARGETS = \
  sql \
  sqlx \
  gorm \
  hello

all: clean fmt lint test build

tools:
	go get golang.org/x/tools/cmd/goimports
	go get golang.org/x/lint/golint

clean:
	rm -rf build/*

fmt:
	gofmt -d . > /tmp/gofmt.log
	test ! -s /tmp/gofmt.log
	rm /tmp/gofmt.log
	goimports -d ./ > /tmp/goimports.log
	test ! -s /tmp/goimports.log
	rm /tmp/goimports.log

fix:
	go fmt ./...
	goimports -w .

lint:
	golint ./... > /tmp/golint.log
	test ! -s /tmp/golint.log
	rm /tmp/golint.log

test:
	go test ./...

build: $(TARGETS)

$(TARGETS):
	go build -o ./build/$@ ./examples/$@/main.go
