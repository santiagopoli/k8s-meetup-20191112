FROM golang:1.12.7 as builder

WORKDIR /build/likes

ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0

COPY go.mod .
COPY go.sum .
RUN go get

COPY main.go .
RUN go build -ldflags="-s -w" -o likes main.go
RUN chmod +x likes

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /build/likes/likes /usr/bin/likes

ENTRYPOINT ["likes"]
