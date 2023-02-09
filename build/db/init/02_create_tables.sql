--ユーザーを切り替え
\c db

--テーブルを作成
CREATE TABLE IF NOT EXISTS users
    (
        "id" bigint NOT NULL,
        "user_id" VARCHAR(36) NOT NULL,
        "first_name" VARCHAR(191) NULL,
        "last_name" VARCHAR(191) NULL,
        "email" VARCHAR(191) NOT NULL UNIQUE,
        "password" VARCHAR(191) NOT NULL
    )
;

--テーブルにデータを挿入
INSERT INTO users VALUES (1, 'user_id_001', 'Hiroshi', 'Iwashita', 'email001@email.com', '00000001');
INSERT INTO users VALUES (2, 'user_id_002', 'Hiromi', 'Koyama', 'email002@email.com', '00000002');

CREATE TABLE IF NOT EXISTS colors
    (
        "id" bigint NOT NULL UNIQUE,
        "name" varchar(30) NOT NULL
    )
;

--テーブルにデータを挿入
INSERT INTO colors VALUES (1, 'Red');
INSERT INTO colors VALUES (2, 'Green');
INSERT INTO colors VALUES (3, 'Blue');