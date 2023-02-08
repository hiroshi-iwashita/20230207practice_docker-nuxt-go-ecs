FROM golang:1.20-bullseye as dev

# Air to for hotreload
RUN go install github.com/cosmtrek/air@v1.29.0

RUN mkdir backend

WORKDIR /backend

ADD . /backend

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]