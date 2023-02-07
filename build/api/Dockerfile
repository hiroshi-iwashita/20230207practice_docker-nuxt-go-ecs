FROM golang:1.19-bullseye as dev

# Air to for hotreload
RUN go install github.com/cosmtrek/air@v1.29.0

# to migrate to mysql
# RUN go install -tags mysql github.com/golang-migrate/migrate/v4/cmd/migrate@latest

RUN mkdir /backend/go/src/app

WORKDIR /backend/go/src/app

ADD . /backend/go/src/app

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]