# syntax=docker/dockerfile:1

## Build
FROM golang:1.16-buster AS build

WORKDIR /app

COPY go.mod ./
 COPY go.sum ./
COPY . .
RUN go mod download

RUN go build -o ./cmd/ ./cmd/community.go

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /app/cmd/community /community

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/community"]
