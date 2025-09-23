FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /app/gitcloser .

FROM gcr.io/distroless/static-debian11

WORKDIR /

COPY --from=builder /app/gitcloser /gitcloser

EXPOSE 8080

ENTRYPOINT ["/gitcloser"]

