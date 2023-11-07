FROM golang:alpine

WORKDIR /app
ADD . .

RUN go build -o cmd/web/main .

EXPOSE 4000

CMD [ "/app/main" ]