FROM golang:1.16.3 as builder

WORKDIR /go/app/book-list

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

COPY . .

RUN cd cmd && go mod download && go build -o ../book-list

FROM alpine:latest
RUN apk --update add --no-cache tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    apk del tzdata

COPY --from=builder /go/app/book-list .

EXPOSE 8000
CMD ["./book-list"]
