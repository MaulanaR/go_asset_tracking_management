FROM golang:1.23 AS builder
RUN mkdir /app
WORKDIR /app
COPY . /app
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/main main.go

FROM gcr.io/distroless/static-debian12
RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/main /app/main
EXPOSE 8080
CMD ["/app/main"]
