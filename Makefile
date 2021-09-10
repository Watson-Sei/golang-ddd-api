make sqlite3/init:
	sqlite3 ddd.db ".databases"
make sqlite3/client:
	sqlite3
make __sqlite3/drop:
	rm -rf ddd.db
make flyway/migrate:
	flyway -configFiles=flyway.conf migrate

make server/up:
	go run cmd/api/main.go