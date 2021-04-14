FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY . .

RUN go get -d -v ./...
RUN go build -o /go/bin ./...

FROM alpine

COPY --from=builder /go/bin/match_gateway /
COPY --from=builder /build/static /static

EXPOSE 8000
ENTRYPOINT ["/match_gateway"]