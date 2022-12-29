# golang-gqlgen-postgresql-example

## Run
```
$ docker-compose up -d #run db
$ make run
```
## Test
```
$ make test
```

## Tree
    - domain/entity: 構造体定義
    - config: 設定
    - graph: gqlgen自動生成(フロントへの返却)
    - io: DBクライアント生成
    - persistence: sqlクエリ(未実装)

## Procedure(gqlgen)
    1. スキーマ定義 → ./graph/schema.graphqlsの修正
    2. gqlコマンド実行 → go gqlgen ./...
    3. ロジック実装 → ./graph/schema.resolvers.goの修正

## Docs
    https://gqlgen.com/getting-started/


## query example
```
query user{
  User(id: "2"){
    id
    name
  }
}
```

```
query offerlist{
  OfferList(id: "2c9h5qymnvt3"){
    offerId
  }
}
```