FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

RUN ls

FROM scratch

COPY --from=builder /dist/match_gateway /

EXPOSE 8000
CMD ["/match_gateway"]