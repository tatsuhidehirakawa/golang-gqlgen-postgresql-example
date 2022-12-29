run:
	go run ./main.go

test:
	go test ./...

boot.dev.back:
	make boot.dev.dbs
	make boot.dev.api

boot.dev.api:
	docker compose --env-file ../dev.env up 125api_dev -d

boot.dev.dbs:
	docker compose --env-file ../dev.env up 134dbs_dev -d

insert.data.api:
	@docker cp ../s_01_src/125api_src/src 125api_dev:/go

eject.data.api:
	@docker cp 125api_dev:/go/src ../s_01_src/125api_src