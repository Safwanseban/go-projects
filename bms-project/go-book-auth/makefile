run:
	go run ./...
proto:
	protoc -I . --go_out . --go_opt=paths=source_relative --go-grpc_out . --go-grpc_opt=paths=source_relative internal/**/pb/*.proto
migrate-up:
	migrate -database ${POSTGRESQL_URL} -path ./pg/migrations up
migrate-down:
	migrate -database ${POSTGRESQL_URL} -path ./pg/migrations down
sqlcgen:
	sqlc generate -f ./pg/sqlc.yaml