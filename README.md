# golnag_architechture


## 以下のコマンドでAPI Schemaを確認できます
```
make gen-swagger
```
URL: http://localhost:8002

## マイグレーション

dry-run(適用されるSQLを確認)
```
make migrate-dry-run
```

apply(実際に適用)
```
make migrate-apply
```