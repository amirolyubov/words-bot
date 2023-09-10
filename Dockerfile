FROM golang:1.21.1 AS builder

RUN go version
RUN apt-get install git
RUN apt install libc6


COPY ./ /what-a-word
WORKDIR /what-a-word

RUN go mod download && go get -u ./...
# RUN CGO_ENABLED=0 go build -o ./app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./main.go 

# second image from first one, but without preinstalled golang 
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /what-a-word/app .
EXPOSE 8080

CMD [ "./app" ]
