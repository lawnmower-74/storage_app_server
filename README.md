## dependencies
```bash
$ go mod tidy
```

## migrate
```bash
$ docker-compose run server go run db/migrate/migrate.go
```