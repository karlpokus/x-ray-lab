# 1. build binary

#FROM golang:alpine AS builder
FROM golang:alpine

WORKDIR /src

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ./hello

EXPOSE 8989

ENTRYPOINT ["/src/hello"]

# 2. copy only binary to new image
#FROM scratch
#WORKDIR /src
#COPY --from=builder /src/hello ./hello
