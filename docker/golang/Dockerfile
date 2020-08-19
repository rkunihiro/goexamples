FROM golang:1.15.0-alpine
ENV CGO_ENABLED=0 \
    GO111MODULE=on
RUN apk add make
RUN go get golang.org/x/tools/cmd/goimports \
 && go get golang.org/x/lint/golint
