genmock:
	mockgen -source="server/services/auth/interface.go" -destination="server/services/auth/mock/auth_mock.go" -package=mock
	mockgen -source="server/services/secret/interface.go" -destination="server/services/secret/mock/secret_mock.go" -package=mock
	mockgen -source="server/storage/interface.go" -destination="server/storage/mock/storage_mock.go" -package=mock
	mockgen -source="client/storage/interface.go" -destination="client/storage/mock/storage_mock.go" -package=mock

migrate:
	go run cmd/server/main.go -d "postgres://postgres:postgres@localhost:5432/tstdb?sslmode=disable" -t "tstdb" -migrate true

srvrun:
	go run cmd/server/main.go -d "postgres://postgres:postgres@localhost:5432/tstdb?sslmode=disable" -t "tstdb" -p ":8085" -s true