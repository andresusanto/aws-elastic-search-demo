FROM golang:1.17-alpine AS builder

RUN apk update \
    && apk add --no-cache git ca-certificates \
    && update-ca-certificates

WORKDIR /app

COPY go.sum go.mod ./
RUN go mod download

COPY . .

ENV GOOS=linux GOARCH=386
RUN go build -ldflags="-w -s" -o /bin/ingester ./cmd/server


FROM alpine:3.13


RUN addgroup appuser && adduser -D -G appuser appuser

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --chown=appuser:appuser --from=builder /bin/ingester /bin/ingester

USER appuser:appuser

EXPOSE 8080

CMD ["/bin/ingester"]
