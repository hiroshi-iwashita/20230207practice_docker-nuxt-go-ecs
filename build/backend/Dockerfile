FROM golang:1.20-bullseye as dev

# Air for hotreload
RUN go install github.com/cosmtrek/air@v1.29.0

RUN mkdir backend

WORKDIR /backend

ADD . /backend

EXPOSE 8080

# air -c [tomlファイル名] // 設定ファイルを指定してair実行(WORKDIRに.air.tomlを配置しておくこと)
CMD ["air", "-c", ".air.toml"]