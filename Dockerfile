FROM golang:1.19-alpine

RUN mkdir -p /go-boiler

WORKDIR /go-boiler
COPY . .

RUN apk update && apk add --no-cache git
RUN go mod tidy

RUN go build /go-boiler

EXPOSE 4000
ENTRYPOINT [ "./go-boiler" ]
