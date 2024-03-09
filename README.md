# Design ER-Diagram 
## https://dbdiagram.io/d/SimpleBank-65e888637570557c71371af2

# Docker
## Setup Postgresql
1. docker ps  =  list container run
2. docker images  =  list images 
3. docker pull  @  https://hub.docker.com/
4. such as pull postgresql 
```zsh
docker pull <name>:<tag>
```
```zsh
docker pull postgresql:15-alpine
```
5. start a container 

```zsh
  docker run --name <container_name> -p <host_port:container_port> -e <environment_variable> -d <image>:<tag>
```

```zsh
  docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine
```

6. or use docker compose (up to me >,<)
```dockerfile
  version: "3.9"
  
  services:
    postgres:
      image: postgres:15-alpine
      container_name: amado_local_db
      # restart: always
      environment:
        POSTGRES_USER: amado_user
        POSTGRES_PASSWORD: dev@1Amado
        POSTGRES_DB: amado_local
      volumes:
        - ./data:/var/lib/postgresql/data
      ports:
        - "5438:5432"
```
7. run command in container
```zsh
docker exec -it <container_name_or_id> <command> [args]
```
```zsh
docker exec -it postgres15 psql -U root
```

simple query
```select now();```

quit `\q`

8. logs container  `docker logs <container_name_or_id>`
9. stop container `docker stop <container_name_or_id>`
10. regardless of their running status `docker ps -a`
11. start again `docker start <container_name_or_id>`
12. shell to use all standard linux commands 
```zsh
docker exec -it postgres15 /bin/sh

pwd

ls -l

createdb --username=root --owner=root simple_bank

psql simple_bank  --> *show simple_bank=#  \q  for exit  

dropdb simple_bank

exit 

or 

docker exec -it postgres15 createdb --username=root --owner=root simple_bank
docker exec -it postgres15 psql -U root simple_bank
docker exec -it postgres15 dropdb simple_bank
```

13. remove container `docker ps -a`
14. show history use command `history | grep "docker run"` other topic besides docker run


credit: https://dev.to/techschoolguru/install-use-docker-postgres-table-plus-to-create-db-schema-44he


# DB Migration
## golang migrate libary
1. golang migrate # https://github.com/golang-migrate/migrate
2. CLI usage # https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
```zsh
brew install golang-migrate
```
check version `migrate -version`

4. add new folder in project `mkdir -p db/migration` for store migration files
check folder `ls -l`

5. create new migration file `migrate -help` for see create command
``` zsh
migrate create -ext sql -dir db/migration -seq init_schema
```

6. copy pg sql form dbdiagram.io paste file up
7. revert change made by the up script on down
``` sql
  DROP TABLE IF EXISTS entries;
  DROP TABLE IF EXISTS transfers;
  DROP TABLE IF EXISTS accounts;
```

8. migration schema file `migrate -help` see detail
```bash
migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up
```

## Create Makefile 
make run command to easy for setup the project on local machine for development
```makefile
postgres:
	docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine

createdb:
	docker exec -it postgres15 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres15 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup migratedown
```

## Compare ORM 
1. GORM - Run slowly on hitg load 
2. SQLX - Quite fast & easy to use, Failure won't occur until runtime
3. SQLC - Very fast & easy


## Installation SQLC
https://sqlc.dev/ 
https://github.com/sqlc-dev/sqlc?tab=readme-ov-file

using command
```bash
brew install sqlc

sqlc version #check version
sqlc help #help

sqlc init #create file sqlc.yaml file

sqlc generate #to generate file 

```

add configuration to sqlc.yaml
```yaml
version: "2"
sql: 
  - schema: "db/migration/"
    queries: "db/query/"
    engine: "postgresql"
    gen:
      go:
        package: db
        out: "db/sqlc"
        emit_json_tags: true
        emit_prepared_queries: false
        emit_interface: false
        emit_exact_table_names: false
```

create query sqlc CRUD Example INSERT
```SQL
-- name: CreateAccount :one
INSERT INTO accounts (
  owner,
  balance,
  currency
) VALUES (
  $1, $2, $3
) RETURNING *;
```
then use command `sqlc generate` 

## GO INIT
```BASH
 go mod init github.com/techschool/simplebank
 go mod tidy
```