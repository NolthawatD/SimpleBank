# Design ER-Diagram 
## https://dbdiagram.io/d/SimpleBank-65e888637570557c71371af2

# Docker
## Setup Postgresql
1. docker ps  =  list container run
2. docker images  =  list images 
3. docker pull  @  https://hub.docker.com/
4. such as pull postgresql 
`docker pull <name>:<tag>`
`docker pull postgresql:15-alpine`
5. start a container 
`docker run --name <container_name> -p <host_port:container_port> -e <environment_variable> -d <image>:<tag>`
`docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine`
6. or use docker compose (up to me >,<)
``` docker-compose.ymal
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
`docker exec -it <container_name_or_id> <command> [args]`
`docker exec -it postgres12 psql -U root`

simple query
`select now();`

quit
`\q`

8. logs container 
`docker logs <container_name_or_id>`

credit: https://dev.to/techschoolguru/install-use-docker-postgres-table-plus-to-create-db-schema-44he


# DB Migration
## golang migrate libary
1. golang migrate # https://github.com/golang-migrate/migrate
2. CLI usage # https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
3. `brew install golang-migrate`
check version `migrate -version`
help `migrate -help`

4. add new folder in project `mkdir -p db/migration` for store migration files
check folder `ls -l`

5. create new migration file `migrate -help` for see create command
`migrate create -ext sql -dir db/migration -seq init_schema`

6. copy pg sql form dbdiagram.io paste file up
7. revert change made by the up script on down
``` file down
  DROP TABLE IF EXISTS entries;
  DROP TABLE IF EXISTS transfers;
  DROP TABLE IF EXISTS accounts;
```


# Compare ORM 
1. GORM - Run slowly on hitg load 
2. SQLX - Quite fast & easy to use, Failure won't occur until runtime
3. SQLC - Very fast & easy