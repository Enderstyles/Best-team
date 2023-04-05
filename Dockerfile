FROM golang:1.19-alpine

RUN apk update && \
    apk add --no-cache git

WORKDIR /app

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["go", "run", "main.go"]