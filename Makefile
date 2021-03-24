# コンテナ起動
up:
	docker-compose
# DBを作成
createdb:
	docker exec -it postgres createdb --username=postgres --owner=postgres simplebank

# DB削除
dropdb:
	docker exec -it postgres dropdb -Upostgres simplebank

# マイグレーション
migrateup:
	migrate -path db/migration -database "postgresql://postgres:secret@localhost:5432/simplebank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:secret@localhost:5432/simplebank?sslmode=disable" -verbose down

# スキーマーの作成
createsc:
	migrate create -ext sql -dir db/migration -seq ${name}

.PHONY: up createdb dropdb migrateup migratedown createsc

# sqlc生成
sqlc:
	sqlc generate

.PHONY: up createdb dropdb migrateup migratedown createsc sqlc