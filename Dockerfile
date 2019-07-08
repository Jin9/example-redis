FROM golang:1.12.6 AS builder
ENV GO111MODULE=on
WORKDIR /src/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -mod=vendor -o goapp server.go

FROM alpine:3.10
RUN apk --no-cache add ca-certificates
COPY --from=builder /src/ .
EXPOSE 8880
ENTRYPOINT ["./goapp"]