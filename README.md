## MemFlash
:zap: Faster access to Memcache &amp; Redis :zap:
-------------
[![GitHub license](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/gunjan5/MemFlash/blob/master/LICENSE)

MemFlash is a multithreaded library with priority based query to access Memcache and Redis cache with Databse faster.
MemFlash also uses gRPC with Protobufs and HTTP/2 to transfer fewer bytes to/from cache/DB with header compression and session reuse (HTTP/2) 

##Tasks:
### Technologies:
- [ ] Memcache 
- [ ] Redis
- [ ] BD? Postgres/mySQL or MongoDB/NoSQL?
- [ ] JSON
- [ ] gRPC
- [ ] Protobufs
- [ ] HTTP/2
- [ ] Performance monitoring
- [ ] Performance graph StatsD?
- [ ] Google AppEngine

### Infrastructure:
- [ ] Statically cross compiled binary 
- [ ] CI/CD with Travis 
- [ ] Docker images
- [ ] Images on DockerHub
- [ ] AWS/GCE



### Setup:
- Install Golang version 1.6 (for native HTTP/2 support)
- Start memcached container: `docker run -p 11211:11211 --name my-memcache -d memcached`
- Start MongoDB container: `docker run -p 27017:27017 -d mongo`
- Start StatsD container: `docker run -d -p 80:80 -p 2003-2004:2003-2004 -p 2023-2024:2023-2024 -p 8125:8125/udp -p 8126:8126 hopsoft/graphite-statsd`

