# -------- build stage --------
FROM golang:1.23 AS build
WORKDIR /app

# copy go module files dulu untuk cache deps
COPY go.mod go.sum ./
RUN go mod download

# copy sisa source
COPY . .

# build static binary (amd64)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -trimpath -ldflags="-s -w" -o server main.go

# -------- run stage --------
FROM gcr.io/distroless/static-debian12
WORKDIR /app
COPY --from=build /app/server .

# vercel akan set PORT; fallback 8080
ENV PORT=8080
EXPOSE 8080

CMD ["/app/server"]
