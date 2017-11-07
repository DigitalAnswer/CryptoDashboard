## DOCKER Container
### mySQL
``` shell
docker run -d --name mysql-dev -e MYSQL_ROOT_PASSWORD=admin --publish 6603:3306 --mount type=bind,source="$(pwd)"/datadir,target=/var/lib/mysql mysql
```