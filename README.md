# CRAWLING

Requirement:

- Crawl [this page](https://malshare.com/daily/) and save to local
- Write some appropriate APIs

## RUN

- Create file .env and update variables according to file example.env
- Go into the root of project and type this command line:

```bash
    go run main.go
```

- Just wait for completely crawling and starting server
- All crawled data will be saved into your database.

## APIs

- Post direct data:

```bash
[POST] /api/push
Eg: curl -X POST "localhost:8080/api/push" -d '{"md5":"hash code md5"}'

Note:
- Make sure that you post correct format hash code.
{
    md5 string
    sha1 string
    sha256 string
    base64 string
    date string
}
```

- Post through raw string:

```javascript
[POST] /api/push/raw-string
Eg: curl -X POST "localhost:8080/api/push/raw-string" -d '{"str":"hello", "date":"2002-10-31"}'

Note:
- Str will be encoded to md5, sha1, sha256.
- If date is omitted, current date will be used.
```

- Get data by ID:

```javascript
[GET] /api/get/{id}
Eg: curl -X GET "localhost:8080/api/get/12345678"
```

- Get data by attributes:

```javascript
[GET] /api/get?date=...&md5=...&sha1=...&sha256=...&base64=...
Eg: curl -X GET "localhost:8080/api/get?md5=laksjdqw"

Note:
- if no queries are provided, return all recodes.
```

- Update data by ID:

```bash
[PATCH] /api/update/{id}

Eg: curl -X PATCH "localhost:8080/api/update/123456" -d '{"md5":"new value"}
```

- Delete data by ID:

```bash
[DELETE] /api/delete/{id}

Eg: curl -X DELETE "localhost:8080/api/delete/123456"
```

- Delete data by attributes:

```bash
[DELETE] /api/delete?date=...&md5=...&sha1=...&sha256=...&base64=...

Eg: curl -X DELETE "localhost:8080/api/delete?md5=asdsaalsdjasda"
```