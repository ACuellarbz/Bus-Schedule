FROM golang:alpine

RUN mkdir /app
WORKDIR /app
RUN apk update && apk add --no-cache git


COPY go.mod go.sum ./

RUN go mod download 

COPY . .

RUN go build -o cmd/web/main .

EXPOSE 4000

CMD [ "/app/main" ]