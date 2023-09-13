FROM okteto/golang:1.21 AS builder

RUN go version

COPY ./ /whataword
WORKDIR /whataword

RUN go mod download && go get -u ./...
RUN CGO_ENABLED=0 go build -o ./app

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /whataword/app .
COPY --from=builder /whataword/fonts/font.ttf fonts/font.ttf

RUN mkdir pic

EXPOSE 8080

CMD [ "./app" ]
