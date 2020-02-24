This section is for docker database volume practice.

We use gPRC-based client & server to do it.

the docker steps below:

- Make sure docker in installed.

1. create a volume for container
```
$ cd ~
$ mkdir data
$ docker volume create s3v6pgdata
$ docker volume ls
```

there's file in /db/data even in new container ( volume works )
```
$ docker run -v s3v6pgdata:/db/data -it ubuntu ls -l /db/data
$ docker run -v s3v6pgdata:/db/data -it ubuntu touch /db/data/file
$ docker run -v s3v6pgdata:/db/data -it ubuntu ls -l /db/data
```

Run the docker command, so that the tar file would be decompressed to docker /volume/
REFERENCE: https://github.com/PacktPublishing/Hands-on-Microservices-with-Go/blob/master/section-3/video-4/data/s3v6.tar.bz2
```
$ docker run -it -v s3v6pgdata:/volume -v /Users/TsengYaoShang/Documents/hands-on-courses/microservices-with-go/section-3-4/data:/backup ubuntu  sh -c "rm -rf /volume/* /volume/..?* /volume/.[!.]* ; tar -C /volume/ -xjf /backup/s3v6.tar.bz2"
```
Then start a postgres container with this volume
```
 $ docker run -v s3v6pgdata:/var/lib/postgresql/data -p 5432:5432 -e POSTGRES_PASSWORD=packt -e POSTGRES_USER=packt -e POSTGRES_DB=wta -d postgres:10.12
```

go to container then try psql console for verifying data
```
$ docker exec -it fcd3a695898f /bin/bash
$ psql -U postgres
#Connect to wta database
\c wta
#Show tables
\dt
#Show schema of players
\d+ players
#Show schema of rankings
\d+ rankings
#Count of players
select count(*) from players;
#Count of Rankins
select count(*) from rankings;
```

Docker run postgresql for usage 
Note: 
1. the port mapping is 5433(local):5432(container)
2. the volume mapping to (container) /var/lib/postgresql/data
```
docker run -v s3v6pgdata:/var/lib/postgresql/data -p 5433:5432 -e POSTGRES_PASSWORD=packt -e POSTGRES_USER=packt -e POSTGRES_DB=wta -d postgres:10.12
```

