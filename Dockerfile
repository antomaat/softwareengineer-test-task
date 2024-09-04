FROM golang:1.23-alpine AS builder

WORKDIR /workspace

COPY go.work ./app/go.mod ./app/go.sum ./

COPY . .

WORKDIR /workspace/app

RUN go mod tidy
RUN go mod download

RUN apk add --no-cache make protobuf protobuf-dev

# Install Go protobuf plugins
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1 \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0

RUN go build -o /workspace/app/bin/app .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /workspace/app/bin/app .

COPY ./app/database.db .

RUN chmod +x ./app

EXPOSE 8080


CMD ["./app"]
