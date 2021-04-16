FROM golang:1.15.7-alpine AS builder
RUN apk add --no-cache git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-X=main.gitCommit=`git rev-parse --short HEAD`" -o bin/k8s-gw

FROM alpine:3
RUN apk --no-cache add ca-certificates
RUN apk add --no-cache curl jq
WORKDIR /root/
COPY --from=builder /app/bin/k8s-gw /usr/local/bin/k8s-gw
ENTRYPOINT ["/usr/local/bin/k8s-gw"]
EXPOSE 4444
