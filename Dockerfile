FROM golang:alpine
WORKDIR /app
COPY . .
RUN apk add git
RUN  go mod download
RUN go mod tidy

RUN go build -o /cmd/web/main.go

EXPOSE 4000

CMD [ "/goddocker" ]