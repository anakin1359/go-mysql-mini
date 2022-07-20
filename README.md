### Setting Environment Variable
---

<br>

module install
```
go get github.com/joho/godotenv@latest
```

```
touch .env .gitignore
```


.env
```
MYSQL_USER=sample_user
MYSQL_PASSWORD=sample_pass
MYSQL_HOST=127.0.0.1
MYSQL_HOST_PORT=1234
MYSQL_DIST_PORT=3306
MYSQL_ROOT_PASSWORD=sample_root_pass
MYSQL_DATABASE=sample_database
```

.gitignore
```
.env
```

<br>

### MySQL DB Container Environment
---

<br>

```
docker-compose -f ./docker-compose.yml up -d
```
```
$ docker ps -a
CONTAINER ID   IMAGE                         COMMAND                  CREATED              STATUS                      PORTS                               NAMES
282fa372a546   mysql:8.0.21                  "docker-entrypoint.s…"   About a minute ago   Up 59 seconds               33060/tcp, 0.0.0.0:3333->3306/tcp   14_mysql_connection_db_1
```
```
$ docker logs -f -t 14_mysql_connection_db_1
2022-07-20T00:12:56.243483900Z 2022-07-20 00:12:56+00:00 [Note] [Entrypoint]: Entrypoint script for MySQL Server 8.0.21-1debian10 started.
2022-07-20T00:12:56.319915200Z 2022-07-20 00:12:56+00:00 [Note] [Entrypoint]: Switching to dedicated user 'mysql'
2022-07-20T00:12:56.324928200Z 2022-07-20 00:12:56+00:00 [Note] [Entrypoint]: Entrypoint script for MySQL Server 8.0.21-1debian10 started.
...
```
```
docker-compose exec db bash
```
```
mysql -u developer -p -h 127.0.0.1 proto
```
```
mysql> show databases;
+--------------------+
| Database           |
+--------------------+
| information_schema |
| proto              |
+--------------------+
2 rows in set (0.00 sec)
```

```
docker-compose down
```

<br>

### Go Run
---

<br>

```
$ go run main.go
Database connection succeeded.
```