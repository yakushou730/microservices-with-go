



```
$ docker volume create usersmariadb
$ docker run -it -v usersmariadb:/volume -v ~/Documents/hands-on-courses/microservices-with-go/section-4-3/data/:/backup ubuntu sh -c "rm -rf /volume/* /volume/..?* /volume/.[!.]*; tar -C /volume/ -xjf /backup/usersmariadb.tar.bz2"
$ docker run -v usersmariadb:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=root-password -p 3307:3306 -d mariadb
$ docker run -p 6380:6379 -d redis

$ docker ps
$ docker exec -it e370465d5935 sh
```

