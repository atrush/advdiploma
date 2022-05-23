mockauth:
	mockgen -source="server/services/auth/interface.go" -destination="server/services/auth/mock/auth_mock.go" -package=mock
mocksecret:
	mockgen -source="server/services/secret/interface.go" -destination="server/services/secret/mock/secret_mock.go" -package=mock
mockstorage:
	mockgen -source="server/storage/interface.go" -destination="server/storage/mock/storage_mock.go" -package=mock

migrate:
	go run cmd/server/main.go -d "postgres://postgres:postgres@localhost:5432/tstdb?sslmode=disable" -t "tstdb" -migrate true