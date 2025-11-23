### How to replicate

Clone repository and generate SQL bindings:

```shell
git clone https://github.com/hreese/issue-demo-sqlc-postgresql-custom-type.git
cd issue-demo-sqlc-postgresql-custom-type
go generate -x ./...
```

Create a new postgresql database and run the binary with the correct [connection string](https://www.postgresql.org/docs/current/libpq-connect.html#LIBPQ-PARAMKEYWORDS):

```shell
DATABASE_URL='postgres://dbuser:dbpass@dbhost:dbport/dbname' go run ./main.go
```