# Build
FROM golang:1.17-alpine AS builder
ARG GITHUB_TOKEN=""
ARG CGO_ENABLED=0
ARG GO111MODULE=on
ARG GOARCH=amd64
ARG GOOS=linux

WORKDIR /gort

COPY go.mod go.sum /gort/
RUN go mod download

COPY . /gort
RUN go build -o app main.go

# Deploy
FROM alpine:3.14 as deployer

COPY --from=builder /gort/app /bin
ENTRYPOINT [ "/bin/app" ]