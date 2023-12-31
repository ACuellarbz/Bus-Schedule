# Build stage
FROM golang:alpine as builder

RUN apk update && apk add --no-cache git
WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download 

COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/web



# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy only the necessary files from the builder stage
COPY --from=builder /app/main .
COPY --from=builder /app/.env . 

EXPOSE 4000

CMD ["./main"]
