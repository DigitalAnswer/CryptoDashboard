## third party dependencies
```go
go get github.com/go-sql-driver/mysql
go get golang.org/x/crypto/bcrypt
```

## DOCKER Container
### mySQL
``` shell
docker run -d --name mysql-dev -e MYSQL_ROOT_PASSWORD=admin --publish 6603:3306 --mount type=bind,source="$(pwd)"/datadir,target=/var/lib/mysql mysql
```

``` sql
DROP TABLE IF EXISTS "currency_info";
CREATE TABLE `currency_info` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `code` varchar(10) NOT NULL DEFAULT '',
  `name` varchar(50) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS "btc_prices";
CREATE TABLE `btc_prices` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `price_usd` double DEFAULT NULL,
  `price_btc` double DEFAULT NULL,
  `PercentChange24h` double DEFAULT NULL,
  `currency_code` varchar(10) NOT NULL DEFAULT '',
  `last_update` int(11) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `time` (`last_update`,`currency_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
```