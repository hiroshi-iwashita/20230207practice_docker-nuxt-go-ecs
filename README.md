## Docker build

### Nuxt ver.3.1.2

- [Nuxt3 × Docker 環境構築（正式版）](https://qiita.com/A-Kira/items/5ce3e1bff34e179ebbc2)

<!-- commandをコメントアウトする -->

$ docker-compose build

$ docker-compose up

$ docker-compose exec frontend bash

$ npx nuxi init . --force

$ npm install

$ npm run dev

$ exit

$ docker-compose down

<!-- commandのコメントアウトを外す -->

$ docker-compose up -d

---

### Golang ver.1.20

% air -c [toml ファイル名] // 設定ファイルを指定して air 実行(WORKDIR に.air.toml を配置しておくこと)

$ docker-compose build

$ docker-compose up

---

### Postgres ver.14.6

- [Docker で PostgreSQL 上にデータベースを作成しよう](https://zenn.dev/farstep/books/7acd1a7fee7e18/viewer/43e8ed)
- [【Docker】postgresql の構築](https://zenn.dev/re24_1986/articles/b76c3fd8f76aec)
- [Docker】PostgreSQL の起動時に初期データをセットアップ](https://amateur-engineer.com/docker-compose-postgresql/)
- [【Docker】起動時に DB の初期データを投入する](https://nandakagoodvibes.hatenablog.com/entry/2022/02/06/163420)
- [golang-migrate での PostgreSQL マイグレーション備忘録](https://zenn.dev/keyamin/articles/24695c455c1591)
- [sqlc と golang-migrate で PostgreSQL を扱う](https://zenn.dev/tchssk/articles/a701d3ce5f9b6b)

$ docker-compose build

$ docker-compose up
