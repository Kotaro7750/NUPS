FROM golang:latest

WORKDIR /app
COPY . .
ENV GO11MODULE=on

RUN go get github.com/pilu/fresh
CMD ["fresh"]