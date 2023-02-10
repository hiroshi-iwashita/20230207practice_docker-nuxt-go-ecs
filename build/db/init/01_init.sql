--ユーザーの作成
CREATE USER root;

--DBの作成
CREATE DATABASE db;

--ユーザーにDBの権限をまとめて付与
GRANT ALL PRIVILEGES ON DATABASE db TO root;
