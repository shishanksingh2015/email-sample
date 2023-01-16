FROM golang:1.18.4-alpine3.15 as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o email-sample server.go

FROM scratch

WORKDIR /app

COPY --from=builder /app/config.yml /app/
COPY --from=builder /app/email-sample /app/
COPY --from=builder /tmp /tmp
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 5000

ENTRYPOINT ["./email-sample"]
