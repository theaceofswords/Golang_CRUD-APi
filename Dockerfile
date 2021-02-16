FROM golang:alpine as builder

LABEL maintainer="Navaneeth.k  <cnavaneeth.k@qburst.com>"

RUN apk update && apk add --no-cache git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download 
COPY . .

RUN go build -o main .

EXPOSE 8080
CMD ["./main"]


#  ++++++++++ try it
# POSTGRES_USER: ${PG_USER:-postgres}
# PGHOST=localhost