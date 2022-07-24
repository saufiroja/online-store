FROM golang:1.18-alpine

RUN apk update && apk upgrade && \
  apk add --no-cache bash git openssh
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main cmd/server.go

EXPOSE 8765 8766
CMD [ "./main" ]