# CRAWLING

Requirement:

- Crawl [this page](https://malshare.com/daily/) and save to local
- Write some appropriate APIs

## RUN

- Create file .env and update variables according to file example.env
- Go into the root of project and type this command line:

```javascript
    go run main.go
```

- Just wait for completely crawling and starting server

## APIs
- Post data to database

```bash
[POST] /api/push/{date}?ext
Eg: curl -X POST http://localhost:8080/api/push/2002-10-31
Note:
 + date must be formatted like yyyy-MM-dd
 + ext: is name of file like: md5,sha1, ...
```

- Get data from database

```bash
[GET] /api/get/{date}?ext
Eg: curl -X GET http://localhost:8080/api/get/2002-10-31
Note:
 + date must be formatted like yyyy-MM-dd
 + ext: is name of file like: md5,sha1, ...
```
