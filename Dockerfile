FROM golang:1.21.1 AS builder

RUN go version
RUN apt-get update
RUN apt-get install -y git
RUN apt-get install -y libc6 
RUN apt-get install -y gcc 
RUN apt-get install -y pkg-config
RUN apt-get install -y libasound2-dev

COPY ./ /what-a-word
WORKDIR /what-a-word

RUN go mod download && go get -u ./...
RUN CGO_ENABLED=0 go build -o ./app

# second image from first one, but without preinstalled golang 
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /what-a-word/app .
EXPOSE 8080

CMD [ "./app" ]
