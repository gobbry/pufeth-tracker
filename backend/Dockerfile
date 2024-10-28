FROM golang:latest AS installer
WORKDIR /app

# Install git for private dependencies if needed
RUN apt-get update && apt-get install -y git && rm -rf /var/lib/apt/lists/*

COPY go.mod ./
COPY go.sum ./

RUN go mod download

FROM installer AS builder
WORKDIR /app
COPY . .

ARG APP_NAME
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main ./apps/${APP_NAME}/main.go

FROM ubuntu:latest AS runner
WORKDIR /

# Install certificates and curl for HTTPS support
RUN apt-get update && \
    apt-get install -y ca-certificates curl && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /app /app
RUN ls

RUN chmod +x /app/main

CMD ["/app/main"]
