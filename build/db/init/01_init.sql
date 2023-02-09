--ユーザーの作成
CREATE USER db;

--DBの作成
CREATE DATABASE db;

--ユーザーにDBの権限をまとめて付与
GRANT ALL PRIVILEGES ON DATABASE db TO db;
