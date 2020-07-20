
FROM golang:1.14

ENV PORT=3000

WORKDIR /go/src/app

COPY . .

RUN go get -u github.com/gofiber/fiber

RUN go get github.com/pilu/fresh

EXPOSE $port

ENTRYPOINT ["fresh"]