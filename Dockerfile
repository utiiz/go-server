FROM golang:1.18.1-alpine3.15 AS builder

WORKDIR /build

COPY main.go .

ENV CGO_ENABLED=0

RUN go mod init github.com/utiiz/go-server

RUN go mod tidy

RUN go build -o go-server main.go


FROM scratch

COPY --from=builder /build/go-server /

ENTRYPOINT ["/go-server"]

EXPOSE 6666
