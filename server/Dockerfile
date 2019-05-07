FROM golang:latest

WORKDIR /app
COPY . .
ENV FO11MODULE=on

RUN go get github.com/pilu/fresh
CMD ["fresh"]