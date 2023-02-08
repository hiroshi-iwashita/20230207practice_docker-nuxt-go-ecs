# Golang 1.20

## go mod

@dir - backend
go mod init [github リポジトリ名など]

## Packages

### Air

- [Docker コンテナで golang をホットリロードする Air を導入](https://zenn.dev/ajapa/articles/bc399c7e4c0def)

@dir - backend

$ touch .air.toml
% # air -c [toml ファイル名] // 設定ファイルを指定して air 実行(WORKDIR に.air.toml を配置する)
