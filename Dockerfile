FROM golang:1.20 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o employee-directory .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/employee-directory /app/employee-directory

WORKDIR /app

CMD ["./employee-directory"]
