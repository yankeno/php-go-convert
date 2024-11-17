# PHP <-> Goの構文変換

PHPとGoの構文を比較する

## 起動手順

```bash
# ローカルのルートディレクトリで実行
docker compose up -d --build
docker compose exec app ash

# コンテナ内で実行
php array.php
go run array_slice.go
```
