



boot.dev.back:
	make boot.dev.dbs
	make boot.dev.api

boot.dev.api:
	docker compose --env-file ../dev.env up 125api_dev -d

boot.dev.dbs:
	docker compose --env-file ../dev.env up 134dbs_dev -d

run:
	go run ./main.go

test:
	go test ./...