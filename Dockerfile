FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . /app

RUN go build -o /wallet-svc

EXPOSE 8080

CMD [ "/wallet-svc" ]